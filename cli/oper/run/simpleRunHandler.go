package run

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
)

func SimpleRunHandler(args []string) {
	if len(args) != 2 {
		panic("Usage: promptc-cli [prompt-file] [input]")
	}
	path := args[0]
	input := args[1]
	txt, err := fetchFile(path)
	if err != nil {
		panic(err)
	}
	file := prompt.ParseFile(txt)
	file.Vars

}

func fetchFile(file string) (txt string, err error) {
	txt, err = iox.ReadAllText(file)
	if err == nil {
		return
	}
	txt, err = iox.ReadAllText(file + ".promptc")
	if err == nil {
		return
	}
	txt, err = iox.ReadAllText(file + ".ptc")
	if err == nil {
		return
	}
	return "", err
}
