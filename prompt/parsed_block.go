package prompt

import (
	"encoding/json"
	"errors"
	"github.com/robertkrimen/otto"
	"strings"
)

type ParsedBlock struct {
	Text    string         `json:"-"`
	VarList []string       `json:"-"`
	Tokens  []BlockToken   `json:"tokens"`
	Extra   map[string]any `json:"extra"`
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
