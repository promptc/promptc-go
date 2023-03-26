package compile

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/prompt"
	"io"
	"os"
	"strings"
)

func AnalyseHandler(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: promptc-cli analyse [prompt-file]")
		return
	}
	promptPath := args[0]
	promptF, err := os.Open(promptPath)
	if err != nil {
		panic(err)
	}
	promptBs, err := io.ReadAll(promptF)
	if err != nil {
		panic(err)
	}

	file := prompt.ParseFile(string(promptBs))
	analyseFile(file)
}

func analyseFile(f *prompt.File) {
	if f == nil {
		shared.ErrorF("File is nil")
		return
	}
	shared.InfoF("Vars in File: ")
	for i, v := range f.ParsedVars {
		fmt.Printf("%s: %#v\n%#v\n", i, v, v.Constraint())
	}
	shared.InfoF("Blocks in File: ")
	for i, b := range f.ParsedPrompt {
		shared.HighlightF("Block %d", i)
		analyse(b)
	}
}

func analyse(p *prompt.ParsedBlock) {
	shared.InfoF("Vars in Prompt Block: ")
	for i, v := range p.VarList {
		fmt.Println(i, v)
	}
	shared.InfoF("Prompt: ")

	varF := console.PrintConfig{
		Foreground: console.Cyan,
		Background: console.None,
		Bold:       true,
	}

	scriptF := console.PrintConfig{
		Foreground: console.Yellow,
		Background: console.None,
		Bold:       true,
		//Underline:  true,
	}
	reserveF := console.PrintConfig{
		Foreground: console.Gray,
		Background: console.None,
		Bold:       true,
		//Underline:  true,
	}

	for _, t := range p.Tokens {
		switch t.Kind {
		case prompt.BlockTokenKindLiter:
			replaced := strings.Replace(t.Text, "{", "{{", -1)
			replaced = strings.Replace(replaced, "}", "}}", -1)
			fmt.Print(replaced)
		case prompt.BlockTokenKindVar:
			varF.Write("{" + t.Text + "}")
		case prompt.BlockTokenKindScript:
			scriptF.Write("{%%\n" + t.Text + "\n%%}")
		case prompt.BlockTokenKindReservedQuota:
			reserveF.Write("'''")
		}
	}
	fmt.Println()
}
