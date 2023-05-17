package prompt

import (
	"encoding/json"
	"errors"
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/prompt/provider"
	"github.com/robertkrimen/otto"
	"strings"
)

type ParsedBlock struct {
	Text    string         `json:"-"`
	VarList []string       `json:"-"`
	Tokens  []BlockToken   `json:"tokens"`
	Extra   map[string]any `json:"extra"`
}

type BlockType string

const (
	PromptBlock BlockType = "prompt"
	RefBlock    BlockType = "ref"
)

func (p *ParsedBlock) Type() BlockType {
	t, ok := p.Extra["type"]
	if !ok {
		return PromptBlock
	}
	switch t {
	case "ref":
		return RefBlock
	case "prompt":
		return PromptBlock
	default:
		return PromptBlock
	}
}

func (p *ParsedBlock) IsRef() bool {
	return p.Type() == RefBlock
}

func (p *ParsedBlock) ToJson() ([]byte, error) {
	return json.Marshal(p)
}

func (p *ParsedBlock) ToMap() map[string]any {
	m := make(map[string]any)
	m["tokens"] = p.Tokens
	m["extra"] = p.Extra
	return m
}
func (p *ParsedBlock) Compile(varMap map[string]string) (compiled string, exceptions []error, fatal bool) {
	return p.CompileWithOption(varMap, true)
}
func (p *ParsedBlock) CompileWithOption(varMap map[string]string, allowScript bool) (compiled string, exceptions []error, fatal bool) {
	fatal = false
	sb := strings.Builder{}
	vm := otto.New()
	for k, v := range varMap {
		err := vm.Set(k, v)
		if err != nil {
			exceptions = append(exceptions, err)
			fatal = true
			return
		}
	}
	//_ = vm.Set("setGlobalVar", func(call otto.FunctionCall) otto.Value {
	//	varName, _ := call.Argument(0).ToString()
	//	varValue, _ := call.Argument(1).ToString()
	//	if varName == "" {
	//		failed, _ := vm.ToValue(false)
	//		return failed
	//	}
	//	varMap[varName] = varValue
	//	ok, _ := vm.ToValue(true)
	//	return ok
	//})

	for _, token := range p.Tokens {
		switch token.Kind {
		case BlockTokenKindLiter:
			sb.WriteString(token.Text)
		case BlockTokenKindVar:
			varVal, ok := varMap[token.Text]
			if !ok {
				exceptions = append(exceptions, errors.New("undefined variable: "+token.Text))
				continue
			}
			sb.WriteString(varVal)
		case BlockTokenKindReservedQuota:
			sb.WriteString("'''")
		case BlockTokenKindScript:
			if !allowScript {
				exceptions = append(exceptions, errors.New("script is not allowed"))
				fatal = true
				return
			}
			script := token.Text
			easyMod := false
			if strings.HasPrefix(script, "E") {
				script = script[1:]
				easyMod = true
			}
			if easyMod {
				script = "result = (function(){\n" + script + "\n})()"
			}
			result, err := vm.Run(script)
			if err != nil {
				exceptions = append(exceptions, err)
				fatal = true
				return
			}
			sb.WriteString(result.String())
		}
	}
	compiled = sb.String()
	return
}

type ReferBlock struct {
	RefTo       string            `json:"ref"`
	VarMap      map[string]string `json:"vars"`
	RefProvider provider.Privider `json:"-"`
}

func (p *ParsedBlock) ToReferBlock(refProvider provider.Privider) (*ReferBlock, error) {
	if !p.IsRef() {
		return nil, errors.New("not a ref block")
	}
	var refBlock ReferBlock
	err := hjson.Unmarshal([]byte(p.Text), &refBlock)
	if err != nil {
		return nil, err
	}
	if refBlock.RefTo == "" {
		return nil, errors.New("no invalid `ref`")
	}
	refBlock.RefProvider = refProvider
	if refBlock.RefProvider == nil {
		return nil, errors.New("no ref provider")
	}
	return &refBlock, nil
}

func (r *ReferBlock) Compile(vars map[string]string) ([]CompiledPrompt, []error) {
	newVars := make(map[string]string)
	for k, v := range vars {
		newVars[k] = v
	}
	for k, v := range r.VarMap {
		if strings.HasPrefix(v, "$") {
			newV := v[1:]
			if strings.HasPrefix(newV, "$") {
				newVars[k] = newV
			} else {
				newVars[k] = vars[newV]
			}
		} else {
			newVars[k] = v
		}
		//fmt.Println(k, "->", v)
	}
	promptTxt, err := r.RefProvider.GetPrompt(r.RefTo)
	if err != nil {
		return nil, []error{err}
	}
	prompt := ParseBasicPrompt(promptTxt)
	prompt.RefProvider = r.RefProvider
	compiledPrompt := prompt.Compile(newVars)

	return compiledPrompt.Prompts, compiledPrompt.Exceptions
}
