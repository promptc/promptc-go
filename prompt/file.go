package prompt

import (
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable"
	"github.com/promptc/promptc-go/variable/interfaces"
)

type Conf struct {
	Model    string `json:"model,omitempty,default=gpt-3.5-turbo"`
	Provider string `json:"provider,omitempty,default=openai"`
}

type FileInfo struct {
	Author  string `json:"author,omitempty"`
	License string `json:"license,omitempty"`
	Project string `json:"project,omitempty"`
	Version string `json:"version,omitempty"`
}

type File struct {
	FileInfo
	Conf         *Conf                          `json:"conf"`
	Prompts      []string                       `json:"prompts"`
	Vars         map[string]string              `json:"vars"`
	ParsedVars   map[string]interfaces.Variable `json:"-"`
	ParsedPrompt []*ParsedBlock                 `json:"-"`
}

var reserved = []string{"conf", "prompts", "vars", "author", "license", "project", "version"}

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

	for _, p := range file.Prompts {
		block := &Block{
			Text: p,
		}
		parsed := block.Parse()
		if parsed == nil {
			fmt.Println("Failed to parse prompt", p)
		}
		for _, v := range parsed.VarList {
			if _, ok := file.Vars[v]; !ok {
				file.Vars[v] = ""
			}
		}
		file.ParsedPrompt = append(file.ParsedPrompt, parsed)
	}

	for k, v := range file.Vars {
		parsed := variable.ParseKeyValue(k, v)
		if parsed == nil {
			fmt.Println("Failed to parse variable", k, v)
		}
		file.ParsedVars[k] = parsed
	}
	return file
}

type CompiledPrompt struct {
	Prompt string
	Extra  map[string]any
}

type CompiledFile struct {
	Info         FileInfo
	Conf         *Conf
	Prompts      []CompiledPrompt
	CompiledVars map[string]string
}

func (f *File) Compile(vars map[string]string) *CompiledFile {
	//varMap := make(map[string]string)
	compiledVars := make(map[string]string)
	for k, v := range f.ParsedVars {
		if val, ok := vars[k]; ok {
			if setted := v.SetValue(val); !setted {
				fmt.Println("Failed to set value", k, val)
				continue
			}
			compiledVars[k] = v.Value()
		}
	}
	var result []CompiledPrompt
	for _, p := range f.ParsedPrompt {
		compiled := p.Compile(compiledVars)
		result = append(result, CompiledPrompt{
			Prompt: compiled,
			Extra:  p.Extra,
		})
	}
	return &CompiledFile{
		Info:         f.FileInfo,
		Conf:         f.Conf,
		Prompts:      result,
		CompiledVars: compiledVars,
	}
}
