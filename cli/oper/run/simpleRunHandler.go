package run

import (
	"fmt"
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
	if len(file.Vars) > 1 {
		fmt.Println("Required following vars:")
		for _, v := range file.Vars {
			fmt.Println("-", v)
		}
		panic("Too many vars")
	}
	varMap := make(map[string]string)
	if len(file.Vars) == 1 {
		for k, _ := range file.Vars {
			varMap[k] = input
		}
	}
	compiled := file.Compile(varMap)
	fmt.Println("Compiled To: ")
	for _, c := range compiled.Prompts {
		fmt.Println("######")
		fmt.Println(c.Prompt)
	}

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
