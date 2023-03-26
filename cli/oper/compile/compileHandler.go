package compile

import (
	"fmt"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/prompt"
	"io"
	"os"
)

func CompileHandler(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: promptc-cli compile [prompt-file] [var-file]")
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

	shared.InfoF("Raw Prompt: ")
	fmt.Println(string(promptBs))

	varMap := shared.IniToMap(string(varBs))
	shared.InfoF("Variables: ")
	for k, v := range varMap {
		fmt.Println(k, ":", v)
	}
	block := &prompt.Block{
		Text: string(promptBs),
	}
	parsed := block.Parse()
	finalPrompt := parsed.Compile(varMap)
	shared.SuccessF("Compiled Prompt: ")
	fmt.Println(finalPrompt)
}
