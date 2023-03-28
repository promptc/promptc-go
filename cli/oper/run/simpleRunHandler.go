package run

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
)

func SimpleRunHandler(args []string) {
	if len(args) < 1 {
		panic("Usage: promptc-cli [prompt-file] [input?]")
	}
	path := args[0]
	input := ""
	if len(args) == 2 {
		input = args[1]
	}
	txt, err := fetchFile(path)
	if err != nil {
		panic(err)
	}
	file := prompt.ParseFile(txt)
	if len(file.Vars) > 1 {
		fmt.Println("Required following vars:")
		for k, v := range file.Vars {
			fmt.Println("-", k, "->", v)
		}
		panic("Too many vars")
	}
	varMap := make(map[string]string)
	if len(file.Vars) == 1 {
		for k, _ := range file.Vars {
			varMap[k] = input
		}
	}
	printSep()
	printInfo(file.FileInfo)
	printSep()
	compiled := file.Compile(varMap)
	fmt.Println("Compiled To: ")
	for _, c := range compiled.Prompts {
		fmt.Println(c.Prompt)
	}

}

func printSep() {
	fmt.Println("================")
}

func printInfo(f prompt.FileInfo) {
	if f.Project != "" {
		fmt.Println("Project:", f.Project)
	}
	if f.Version != "" {
		fmt.Println("Version:", f.Version)
	}
	if f.Author != "" {
		fmt.Println("Author:", f.Author)
	}
	if f.License != "" {
		fmt.Println("License:", f.License)
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
