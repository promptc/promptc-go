package prompt

import (
	"github.com/promptc/promptc-go/prompt/provider"
	"github.com/promptc/promptc-go/variable/interfaces"
	"github.com/sashabaranov/go-openai"
)

type DriverExtra struct {
	Temperature *float32 `json:"temperature,omitempty"`
	Stop        []string `json:"stop,omitempty"`
}

type Conf struct {
	DriverExtra
	Model    string `json:"model,omitempty,default=gpt-3.5-turbo"`
	Provider string `json:"provider,omitempty,default=openai"`
}

type FileInfo struct {
	Author  string `json:"author,omitempty"`
	License string `json:"license,omitempty"`
	Project string `json:"project,omitempty"`
	Version string `json:"version,omitempty"`
}

type PromptC struct {
	FileInfo
	Conf          *Conf                          `json:"conf,omitempty"`
	Vars          map[string]string              `json:"vars"`
	Prompts       []string                       `json:"prompts"`
	VarConstraint map[string]interfaces.Variable `json:"-"`
	ParsedPrompt  []*ParsedBlock                 `json:"-"`
	Exceptions    []error                        `json:"-"`
	RefProvider   provider.Privider              `json:"-"`
}

func (f *PromptC) GetConf() Conf {
	if f.Conf == nil {
		return Conf{
			Provider: "openai",
			Model:    "gpt-3.5-turbo",
		}
	}
	return *f.Conf
}

var reserved = []string{"conf", "prompts", "vars", "author", "license", "project", "version"}

type CompiledPrompt struct {
	Prompt string
	Extra  map[string]any
}

type CompiledPromptC struct {
	Fatal        bool
	Info         FileInfo
	Conf         *Conf
	Prompts      []CompiledPrompt
	CompiledVars map[string]string
	Exceptions   []error
}

func ReservedKeys() []string {
	return reserved
}

func (c *CompiledPromptC) OpenAIChatCompletionMessages(ignoreEmptyPrompt bool) []openai.ChatCompletionMessage {
	var messages []openai.ChatCompletionMessage
	for _, _p := range c.Prompts {
		if ignoreEmptyPrompt && _p.Prompt == "" {
			continue
		}

		role := "user"
		if len(_p.Extra) > 0 {
			ok := false
			var a any
			a, ok = _p.Extra["role"]
			if ok {
				role, ok = a.(string)
				if !ok {
					role = "user"
				}
			}
		}

		message := openai.ChatCompletionMessage{
			Role:    role,
			Content: _p.Prompt,
		}
		messages = append(messages, message)
	}
	return messages
}
