package fmtt

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	ptProvider "github.com/promptc/promptc-go/prompt/provider"
	"path/filepath"
)

func FormatHandler(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: promptc fmt <input> <output>")
		return
	}
	input := args[0]
	output := args[1]

	fileStr, err := iox.ReadAllText(input)
	if err != nil {
		panic(err)
	}
	file := prompt.ParseFile(fileStr)
	file.RefProvider = &ptProvider.FileProvider{
		BasePath: filepath.Dir(input),
	}
	str := file.Formatted()
	err = iox.WriteAllText(output, str)
}
