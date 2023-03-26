package prompt

import (
	"github.com/robertkrimen/otto"
	"strings"
)

type ParsedBlock struct {
	text    string
	varList []string
	tokens  []BlockToken
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
	for _, token := range p.tokens {
		switch token.Kind {
		case BlockTokenKindLiter:
			sb.WriteString(token.Text)
		case BlockTokenKindVar:
			sb.WriteString(varMap[token.Text])
		case BlockTokenKindScript:
			result, err := vm.Run(token.Text)
			if err != nil {
				panic(err)
			}
			sb.WriteString(result.String())
		}
	}
	return sb.String()
}
