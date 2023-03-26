package prompt

import (
	"strings"
)

type Block struct {
	text string
}

// Text := Alphabets | {{ | }}
// Var := {Text}

func (b *Block) Parse() *ParsedBlock {
	rs := []rune(b.text)
	var varList []string
	var tokens []BlockToken
	isOpen := false
	sb := strings.Builder{}
	var prev rune = 0
	for _, r := range rs {
		if r == '{' {
			// {{ <- 变身为 {
			if prev == '{' {
				sb.WriteRune('{')
				isOpen = false
				r = 0
				goto nextStep
			}
			// 开始解析变量
			isOpen = true
			goto nextStep
		}
		if r == '}' {
			// 如果是 }} 则变身为 }
			if isOpen {
				isOpen = false
				name := strings.TrimSpace(sb.String())
				sb.Reset()
				varList = append(varList, name)
				tokens = append(tokens, BlockToken{name, BlockTokenKindVar})
				r = 0
				goto nextStep
			}
			if prev == '}' {
				sb.WriteRune('}')
				r = 0
				goto nextStep
			}
			goto nextStep

		}
		if isOpen && prev == '{' {
			val := sb.String()
			tokens = append(tokens, BlockToken{val, BlockTokenKindLiter})
			sb.Reset()
		}
		if !isOpen && prev == '}' {
			sb.WriteRune('}')
		}
		sb.WriteRune(r)
	nextStep:
		//fmt.Printf("r=%s, prev=%s, isOpen=%#v, sb=%#v\n", sr(r), sr(prev), isOpen, sb.String())
		prev = r
	}
	if sb.Len() > 0 {
		kind := BlockTokenKindLiter
		if isOpen {
			kind = BlockTokenKindVar
		}
		tokens = append(tokens, BlockToken{sb.String(), kind})
	}
	return &ParsedBlock{
		text:    b.text,
		varList: varList,
		tokens:  tokens,
	}
}

func sr(r rune) string {
	if r == 0 {
		return " "
	}
	if r == '\n' {
		return "\\n"
	}
	return string(r)
}

type ParsedBlock struct {
	text    string
	varList []string
	tokens  []BlockToken
}
