package utils

import "github.com/hjson/hjson-go/v4"

func Hjson(a any) string {
	bs, err := hjson.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func HjsonNoIdent(a any) string {
	if a == nil {
		return ""
	}
	bs, err := hjson.MarshalWithOptions(a, hjson.EncoderOptions{
		IndentBy:       "",
		EmitRootBraces: false,
	})
	if err != nil {
		return err.Error()
	}
	return string(bs)
}

func HjsonNoBrace(a any) string {
	if a == nil {
		return ""
	}
	option := hjson.DefaultOptions()
	option.EmitRootBraces = false
	bs, err := hjson.MarshalWithOptions(a, option)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
