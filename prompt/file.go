package prompt

import (
	"fmt"

	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable"
	"github.com/promptc/promptc-go/variable/interfaces"
)

func ParsePromptC(content string) *PromptC {
	file := &PromptC{
		VarConstraint: make(map[string]interfaces.Variable),
		Vars:          make(map[string]string),
	}
	var fileM map[string]any
	var hjsonResult any
	err := hjson.Unmarshal([]byte(content), &hjsonResult)
	if err != nil {
		panic(err)
	}
	fileM, ok := hjsonResult.(map[string]any)
	if !ok {
		fileM = make(map[string]any)
		file.Prompts = append(file.Prompts, content)
		file.Conf = &Conf{
			Model:    "gpt-3.5-turbo",
			Provider: "openai",
		}
	} else {
		err = hjson.Unmarshal([]byte(content), file)
		if err != nil {
			panic(err)
		}
	}

	// remove reserved keys
	for _, r := range reserved {
		delete(fileM, r)
	}

	// parse wild defined vars
	for k, v := range fileM {
		result, ok := v.(string)
		if !ok {
			continue
		}
		file.Vars[k] = result
	}

	file.parsePrompt()
	file.parseVariable()
	return file
}

func (f *PromptC) parseVariable() {
	for k, v := range f.Vars {
		parsed, err := variable.ParseKeyValue(k, v)
		if parsed != nil {
			f.VarConstraint[k] = parsed
		}
		if err != nil {
			f.Exceptions = append(f.Exceptions, fmt.Errorf("failed to parse variable %s -> %s", k, v))
			f.Exceptions = append(f.Exceptions, err)
		}
	}
}

func (f *PromptC) parsePrompt() {
	for promptId, p := range f.Prompts {
		block := &Block{
			Text: p,
		}
		parsed := block.Parse()
		if parsed == nil {
			f.Exceptions = append(f.Exceptions, fmt.Errorf("failed to parse prompt %d", promptId))
		}
		for _, v := range parsed.VarList {
			if _, ok := f.Vars[v]; !ok {
				f.Vars[v] = ""
			}
		}

		f.ParsedPrompt = append(f.ParsedPrompt, parsed)
	}
}

func ParseBasicPrompt(content string) *PromptC {
	file := &PromptC{
		VarConstraint: make(map[string]interfaces.Variable),
		Vars:          make(map[string]string),
		Prompts:       []string{content},
		Conf: &Conf{
			Model:    "gpt-3.5-turbo",
			Provider: "openai",
		},
	}

	file.parsePrompt()
	file.parseVariable()
	return file
}

func (f *PromptC) Compile(vars map[string]string) *CompiledPromptC {
	//varMap := make(map[string]string)
	fileFatal := false
	compiledVars := make(map[string]string)
	var errs []error
	errs = append(errs, f.Exceptions...)
	for k, v := range f.VarConstraint {
		if val, ok := vars[k]; ok {
			if setted := v.SetValue(val); !setted {
				errs = append(errs, fmt.Errorf("failed to set value %s %s", k, val))
				continue
			}
			compiledVars[k] = v.Value()
		}
	}
	for k, v := range vars {
		if _, ok := compiledVars[k]; !ok {
			compiledVars[k] = v
		}
	}
	var result []CompiledPrompt
	for _, p := range f.ParsedPrompt {
		if p.IsRef() {
			refB, _err := p.ToReferBlock(f.RefProvider)
			if _err != nil {
				errs = append(errs, _err)
				continue
			}
			compiled, _errs := refB.Compile(compiledVars)
			if _err != nil {
				errs = append(errs, _errs...)
				continue
			}
			result = append(result, compiled...)
			continue
		}
		compiled, exp, fatal := p.Compile(compiledVars)
		if len(exp) > 0 {
			errs = append(errs, exp...)
		}
		if fatal {
			fileFatal = true
			goto compiled
		}
		result = append(result, CompiledPrompt{
			Prompt: compiled,
			Extra:  p.Extra,
		})
	}
compiled:
	return &CompiledPromptC{
		Fatal:        fileFatal,
		Info:         f.FileInfo,
		Conf:         f.Conf,
		Prompts:      result,
		CompiledVars: compiledVars,
		Exceptions:   errs,
	}
}
