package compile

import (
	"fmt"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/prompt"
	"github.com/promptc/promptc-go/prompt/provider"
	"github.com/promptc/promptc-go/utils"
	"io"
	"os"
	"path/filepath"
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
	shared.InfoF("Entered Variables: ")
	for k, v := range varMap {
		fmt.Println(k, ":", v)
	}

	file := prompt.ParsePromptC(string(promptBs))
	file.RefProvider = &provider.FileProvider{
		BasePath: filepath.Dir(promptPath),
	}
	shared.InfoF("Prompt Conf: ")
	fmt.Println(utils.Hjson(file.Conf))

	shared.InfoF("Compiling...")
	compiled := file.Compile(varMap)
	shared.InfoF("Compiled Vars: ")
	for k, v := range compiled.CompiledVars {
		fmt.Println(k, ":", v)
	}
	shared.InfoF("Compiled Prompt: ")
	for _, c := range compiled.Prompts {
		shared.InfoF("Extra:")
		for k, v := range c.Extra {
			fmt.Println(k, ":", v)
		}
		shared.InfoF("Prompt:")
		fmt.Println(c.Prompt)
	}
	shared.ErrorF("Compiled Exceptions: ")
	for _, e := range compiled.Exceptions {
		fmt.Println(e)
	}
}
