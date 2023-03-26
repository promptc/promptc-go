package prompt

import (
	"fmt"
	"testing"
)

func TestParseBlock(t *testing.T) {
	text := `
You Entered: {x}
Prompt Compiled: {%
	if (x == "1") {
		result = "Hello"
	} else {
		result = "Word!";
	}
%}
`
	varMap := map[string]string{
		"x": "KevinZonda",
	}
	b := &Block{
		text: text,
	}
	parsed := b.Parse()
	rst := parsed.Compile(varMap)
	fmt.Println(rst)

	varMap["x"] = "1"
	rst = parsed.Compile(varMap)
	fmt.Println(rst)

}

func TestParseBlockEasy(t *testing.T) {
	text := `
You Entered: {x}
Prompt Compiled: {%E
	if (x == "1") {
		return "Hello"
	} else {
		return "Word!";
	}
%}
`
	varMap := map[string]string{
		"x": "KevinZonda",
	}
	b := &Block{
		text: text,
	}
	parsed := b.Parse()
	rst := parsed.Compile(varMap)
	fmt.Println(rst)

	varMap["x"] = "1"
	rst = parsed.Compile(varMap)
	fmt.Println(rst)

}
