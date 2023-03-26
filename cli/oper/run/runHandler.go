package run

import (
	"fmt"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/prompt"
	"io"
	"os"
)

func RunHandler(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: promptc-cli run [prompt-file] [var-file]")
		return
	}
	promptPath := args[0]
	varPath := args[1]
	promptF, err := os.Open(promptPath)
	if err != nil {
		panic(err)
	}
	promptBs, err := io.ReadAll(promptF)
	if err != nil {
		panic(err)
	}
	varF, err := os.Open(varPath)
	if err != nil {
		panic(err)
	}
	varBs, err := io.ReadAll(varF)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(varBs))
	// fmt.Println(string(promptBs))

	varMap := shared.IniToMap(string(varBs))
	block := &prompt.Block{
		Text: string(promptBs),
	}
	parsed := block.Parse()
	finalPrompt := parsed.Compile(varMap)
	fmt.Println(finalPrompt)
}
