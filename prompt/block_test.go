package prompt

import (
	"fmt"
	"strings"
	"testing"
)

func TestBlockParser(t *testing.T) {
	text := `
This is a block of {Text}, Ha
pp Day {x
}
MM {{}} {{{x}}}
`
	block := Block{
		text: text,
	}
	parsed := block.Parse()
	printBlock(parsed)
}

func printTokens(tokens []BlockToken) {
	for _, t := range tokens {
		fmt.Printf("%#v\n", t)
	}
}

func printBlock(b *ParsedBlock) {
	fmt.Printf("TEXT : %#v\n", b.text)
	fmt.Printf("VAR  : %#v\n", b.varList)
	fmt.Printf("TOKEN:\n")
	printTokens(b.tokens)
	fmt.Printf("BACK : %#v\n", backToText(b.tokens))
}

func backToText(tokens []BlockToken) string {
	sb := strings.Builder{}
	for _, t := range tokens {
		if t.Kind != BlockTokenKindVar {
			replaced := strings.ReplaceAll(t.Text, "{", "{{")
			replaced = strings.ReplaceAll(replaced, "}", "}}")
			sb.WriteString(replaced)
			continue
		}

		sb.WriteString("{")
		sb.WriteString(t.Text)
		sb.WriteString("}")
	}
	return sb.String()
}
