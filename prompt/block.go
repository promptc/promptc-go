package prompt

import (
	"strings"
)

type Block struct {
	Text string
}

// Text := Alphabets | {{ | }}
// Var := {Text}

func (b *Block) Parse() *ParsedBlock {
	rs := []rune(b.Text)
	var varList []string
	var tokens []BlockToken
	isOpen := false
	isScriptOpen := false
	sb := strings.Builder{}
	var prev rune = 0
	for _, r := range rs {
		if r == '{' {
			if isScriptOpen {
				sb.WriteRune(r)
				goto nextStep
			}
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
		if r == '%' {
			if prev == '{' && isOpen && !isScriptOpen {
				isScriptOpen = true
				if sb.Len() > 0 {
					val := sb.String()
					tokens = append(tokens, BlockToken{val, BlockTokenKindLiter})
					sb.Reset()
				}
			} else {
				sb.WriteRune(r)
			}
			goto nextStep

		}
		if r == '}' {
			// 如果是 }} 则变身为 }
			if isOpen {
				kind := BlockTokenKindVar
				if isScriptOpen {
					if prev == '%' {
						isScriptOpen = false
						kind = BlockTokenKindScript
					} else {
						sb.WriteRune('}')
						r = 0
						goto nextStep
					}
				}
				isOpen = false
				name := strings.TrimSpace(sb.String())
				if kind == BlockTokenKindScript {
					name = strings.Trim(name, "%")
					name = strings.TrimSpace(name)
				}
				sb.Reset()
				varList = append(varList, name)
				tokens = append(tokens, BlockToken{name, kind})
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
		if isOpen && prev == '{' && !isScriptOpen {
			val := sb.String()
			tokens = append(tokens, BlockToken{val, BlockTokenKindLiter})
			sb.Reset()
		}
		if !isOpen && prev == '}' {
			sb.WriteRune('}')
		}
		sb.WriteRune(r)
	nextStep:
		//fmt.Printf("r=%s, prev=%s, isOpen=%#v, isScript=%#v, sb=%#v\n", sr(r), sr(prev), isOpen, isScriptOpen, sb.String())
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
		Text:    b.Text,
		VarList: varList,
		Tokens:  tokens,
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
