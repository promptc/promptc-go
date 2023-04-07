package types

import "github.com/hjson/hjson-go/v4"

func hjsonNoIndent(a any) string {
	if a == nil {
		return ""
	}
	option := hjson.DefaultOptions()
	option.EmitRootBraces = false
	option.Eol = ", "
	option.IndentBy = ""
	option.BracesSameLine = true

	bs, err := hjson.MarshalWithOptions(a, option)
	if err != nil {
		return err.Error()
	}
	return " { " + string(bs) + " }"
}
