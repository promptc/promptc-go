package compile

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/prompt"
)

// AnalyzeHandler handles the analysis of a prompt file.
//
// args: a slice of strings representing the command line arguments.
// It should contain only one element, which is the path to the prompt file.
//
// Returns nothing.
func AnalyseHandler(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: promptc-cli analyze [prompt-file]")
		return
	}

	promptPath := args[0]
	promptFile, err := os.Open(promptPath)
	if err != nil {
		panic(err)
	}
	defer promptFile.Close()

	promptContent, err := io.ReadAll(promptFile)
	if err != nil {
		panic(err)
	}

	file := prompt.ParsePromptC(string(promptContent))
	analyseFile(file)
}

func analyseFile(f *prompt.PromptC) {
	if f == nil {
		shared.ErrorF("File is nil")
		return
	}
	shared.InfoF("Vars in File: ")
	for i, v := range f.VarConstraint {
		fmt.Printf("%s: %#v\n%#v\n", i, v, v.Constraint())
	}
	shared.InfoF("Blocks in File: ")
	for i, b := range f.ParsedPrompt {
		shared.HighlightF("Block %d", i)
		analyse(b)
	}
}

// analyse processes a Prompt block and prints information about its tokens,
// variables, and extra values.
//
// p: a pointer to a Prompt.ParsedBlock struct
func analyse(p *prompt.ParsedBlock) {
	shared.InfoF("Vars in Prompt Block: ")
	for i, v := range p.VarList {
		fmt.Println(i, v)
	}

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
	shared.InfoF("Extra: ")
	for k, v := range p.Extra {
		fmt.Printf("%s: %#v\n", k, v)
	}

	shared.InfoF("Tokens: ")
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
