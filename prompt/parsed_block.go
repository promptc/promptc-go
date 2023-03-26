package prompt

import (
	"github.com/robertkrimen/otto"
	"strings"
)

type ParsedBlock struct {
	Text    string
	VarList []string
	Tokens  []BlockToken
}

func (p *ParsedBlock) Compile(varMap map[string]string) string {
	sb := strings.Builder{}
	vm := otto.New()
	for k, v := range varMap {
		err := vm.Set(k, v)
		if err != nil {
			panic(err)
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
			sb.WriteString(varMap[token.Text])
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
				panic(err)
			}
			sb.WriteString(result.String())
		}
	}
	return sb.String()
}
