package prompt

import "github.com/promptc/promptc-go/variable/interfaces"

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
	Conf          *Conf                          `json:"conf"`
	Prompts       []string                       `json:"prompts"`
	Vars          map[string]string              `json:"vars"`
	VarConstraint map[string]interfaces.Variable `json:"-"`
	ParsedPrompt  []*ParsedBlock                 `json:"-"`
	Exceptions    []error                        `json:"-"`
}

var reserved = []string{"conf", "prompts", "vars", "author", "license", "project", "version"}

type CompiledPrompt struct {
	Prompt string
	Extra  map[string]any
}

type CompiledFile struct {
	Fatal        bool
	Info         FileInfo
	Conf         *Conf
	Prompts      []CompiledPrompt
	CompiledVars map[string]string
	Exceptions   []error
}
