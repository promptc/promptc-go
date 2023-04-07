package types

import "github.com/hjson/hjson-go/v4"

func hjsonNoIndent(a any) string {
	if a == nil {
		return ""
	}
	bs, err := hjson.MarshalWithOptions(a, hjson.EncoderOptions{
		IndentBy:       "",
		EmitRootBraces: true,
	})
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
