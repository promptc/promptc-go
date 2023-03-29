package prompt

import (
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable"
	"github.com/promptc/promptc-go/variable/interfaces"
)

func ParseFile(content string) *File {
	file := &File{
		ParsedVars: make(map[string]interfaces.Variable),
		Vars:       make(map[string]string),
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

func (f *File) parseVariable() {
	for k, v := range f.Vars {
		parsed := variable.ParseKeyValue(k, v)
		if parsed == nil {
			f.Exceptions = append(f.Exceptions, fmt.Errorf("failed to parse variable %s -> %s", k, v))
			continue
		}
		f.ParsedVars[k] = parsed
	}
}

func (f *File) parsePrompt() {
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

func ParseUnstructuredFile(content string) *File {
	file := &File{
		ParsedVars: make(map[string]interfaces.Variable),
		Vars:       make(map[string]string),
		Prompts:    []string{content},
		Conf: &Conf{
			Model:    "gpt-3.5-turbo",
			Provider: "openai",
		},
	}

	file.parsePrompt()
	file.parseVariable()
	return file
}

func (f *File) Compile(vars map[string]string) *CompiledFile {
	//varMap := make(map[string]string)
	fileFatal := false
	compiledVars := make(map[string]string)
	var errs []error
	errs = append(errs, f.Exceptions...)
	for k, v := range f.ParsedVars {
		if val, ok := vars[k]; ok {
			if setted := v.SetValue(val); !setted {
				errs = append(errs, fmt.Errorf("failed to set value %s %s", k, val))
				continue
			}
			compiledVars[k] = v.Value()
		}
	}
	var result []CompiledPrompt
	for _, p := range f.ParsedPrompt {
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
	return &CompiledFile{
		Fatal:        fileFatal,
		Info:         f.FileInfo,
		Conf:         f.Conf,
		Prompts:      result,
		CompiledVars: compiledVars,
		Exceptions:   errs,
	}
}
