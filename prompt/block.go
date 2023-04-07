package prompt

import (
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/prompt/provider"
	"strings"
)

type Block struct {
	Text        string
	RefProvider provider.Privider
}

// Text := Alphabets | {{ | }}
// Var := {Text}

func (b *Block) Parse() *ParsedBlock {
	lines := strings.SplitN(b.Text, "\n", 2)
	firstLine := ""
	toParse := ""
	if len(lines) == 1 {
		toParse = b.Text
	} else {
		firstLine = strings.TrimSpace(lines[0])
		toParse = lines[1]
	}
	extra := make(map[string]any)
	if len(firstLine) > 0 && firstLine != "{}" {
		err := hjson.Unmarshal([]byte(firstLine), &extra)
		if err != nil {
			toParse = b.Text
		}
	}
	rs := []rune(toParse)
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
				sb.Reset()
				if kind == BlockTokenKindScript {
					name = strings.Trim(name, "%")
					name = strings.TrimSpace(name)
					if name == "Q" {
						kind = BlockTokenKindReservedQuota
					}
				} else {
					varList = append(varList, name)
				}

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
		name := sb.String()
		sb.Reset()
		if isOpen {
			kind = BlockTokenKindVar
			name = strings.TrimSpace(name)
			if isScriptOpen {
				name = strings.Trim(name, "%")
				name = strings.TrimSpace(name)
				kind = BlockTokenKindScript
				if name == "Q" {
					kind = BlockTokenKindReservedQuota
				}
			}
		}
		tokens = append(tokens, BlockToken{name, kind})
	}
	return &ParsedBlock{
		Text:        b.Text,
		VarList:     varList,
		Tokens:      tokens,
		Extra:       extra,
		RefProvider: b.RefProvider,
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
