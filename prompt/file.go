package prompt

import (
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/variable"
	"github.com/promptc/promptc-go/variable/interfaces"
)

type Conf struct {
	Model    string `json:"model,omitempty,default=gpt-3.5-turbo"`
	Provider string `json:"provider,omitempty,default=OpenAI"`
}

type File struct {
	Conf         *Conf                          `json:"conf"`
	Prompts      []string                       `json:"prompts"`
	Vars         map[string]string              `json:"vars"`
	ParsedVars   map[string]interfaces.Variable `json:"-"`
	ParsedPrompt []*ParsedBlock                 `json:"-"`
}

var reserved = []string{"conf", "prompts", "vars"}

func ParseFile(content string) *File {
	file := &File{
		ParsedVars: make(map[string]interfaces.Variable),
		Vars:       make(map[string]string),
	}
	fileM := make(map[string]any)
	err := hjson.Unmarshal([]byte(content), file)
	if err != nil {
		panic(err)
	}
	err = hjson.Unmarshal([]byte(content), &fileM)
	if err != nil {
		panic(err)
	}
	for _, r := range reserved {
		delete(fileM, r)
	}
	for k, v := range fileM {
		result, ok := v.(string)
		if !ok {
			continue
		}
		file.Vars[k] = result
	}

	for k, v := range file.Vars {
		parsed := variable.ParseKeyValue(k, v)
		if parsed == nil {
			fmt.Println("Failed to parse variable", k, v)
		}
		file.ParsedVars[k] = parsed
	}

	for _, p := range file.Prompts {
		block := &Block{
			Text: p,
		}
		parsed := block.Parse()
		if parsed == nil {
			fmt.Println("Failed to parse prompt", p)
		}
		file.ParsedPrompt = append(file.ParsedPrompt, parsed)
	}
	return file
}