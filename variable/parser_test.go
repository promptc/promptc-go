package variable

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	v := "x : int { min: 1, max: 10 }"
	fmt.Printf("ToDo: %#v\n", v)
	p := Parse(v)
	if p == nil {
		t.Error("Failed to parse variable")
	}
	fmt.Printf("Type: %#v\n", p)
	fmt.Printf("Cons: %#v\n", p.Constraint())
	fmt.Printf("Name: %#v\n", p.Name())
	fmt.Printf("Type: %#v\n", p.Type())
	fmt.Printf("SetV: %#v\n", p.SetValue("5"))
	fmt.Printf("GetV: %#v\n", p.Value())
	fmt.Printf("HasV: %#v\n", p.HasValue())
}
