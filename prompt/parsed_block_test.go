package prompt

import (
	"encoding/json"
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
		Text: text,
	}
	parsed := b.Parse()
	rst, exp, fatal := parsed.Compile(varMap)
	fmt.Println(rst)
	fmt.Println(exp)
	fmt.Println(fatal)

	varMap["x"] = "1"
	rst, exp, fatal = parsed.Compile(varMap)
	fmt.Println(rst)
	fmt.Println(exp)
	fmt.Println(fatal)

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
		Text: text,
	}
	parsed := b.Parse()
	jsoned, err := parsed.ToJson()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsoned))
	maped := parsed.ToMap()
	Njson, err := json.Marshal(maped)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(Njson))

	rst, exp, fatal := parsed.Compile(varMap)
	fmt.Println(rst)
	fmt.Println(exp)
	fmt.Println(fatal)

	varMap["x"] = "1"
	rst, exp, fatal = parsed.Compile(varMap)
	fmt.Println(rst)
	fmt.Println(exp)
	fmt.Println(fatal)

}
