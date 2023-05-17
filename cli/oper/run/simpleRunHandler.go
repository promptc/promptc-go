package run

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
	ptProvider "github.com/promptc/promptc-go/prompt/provider"
	"github.com/promptc/promptc-go/utils"
)

// SimpleRunHandler handles a command line prompt with input variables.
//
// args: a slice of string arguments passed to the program.
// It expects at least one argument, the path to a promptc file.
//
// Returns nothing.
func SimpleRunHandler(args []string) {
	if len(args) < 1 {
		panic("Usage: promptc-cli [prompt-file] [input?]")
	}

	path := args[0]
	inputs := args[1:]
	var input string
	if len(args) == 2 {
		input = args[1]
	}

	txt, structured, err := fetchFile(path)
	if err != nil {
		panic(err)
	}

	var file *prompt.PromptC
	if structured {
		file = prompt.ParsePromptC(txt)
	} else {
		file = prompt.ParseBasicPrompt(txt)
	}
	file.RefProvider = &ptProvider.FileProvider{
		BasePath: filepath.Dir(path),
	}

	varMap := make(map[string]string)
	if len(file.Vars) > 1 || (len(inputs) == 0 && len(file.Vars) > 0) {
		fmt.Println("Please enter the following variables:")
		for k, v := range file.Vars {
			fmt.Print(k, " (", v, "): ")
			input, err := console.ReadLine()
			if err != nil {
				panic(err)
			}
			varMap[k] = input
		}
	}

	if len(file.Vars) == 1 {
		for k := range file.Vars {
			varMap[k] = input
		}
	} else {
		if len(file.ParsedPrompt) == 0 {
			panic("No prompts")
		}
		lastBlock := file.ParsedPrompt[len(file.ParsedPrompt)-1]
		lastBlock.Tokens = append(lastBlock.Tokens, prompt.BlockToken{
			Kind: prompt.BlockTokenKindLiter,
			Text: " " + strings.Join(inputs, " "),
		})
		file.Prompts[len(file.Prompts)-1] += " " + strings.Join(inputs, " ")
	}

	printSep()
	printInfo(file.FileInfo)
	printSep()

	compiled := file.Compile(varMap)

	fmt.Println("Compiled to:")
	for i, c := range compiled.Prompts {
		shared.InfoF("Prompt #%d [%s]: ", i, utils.HjsonNoIdent(c.Extra))
		fmt.Println(c.Prompt)
	}

	if len(compiled.Exceptions) > 0 {
		printSep()
		shared.ErrorF("Compiled Exceptions: ")
		for _, e := range compiled.Exceptions {
			fmt.Println(e)
		}
	}

	provider := strings.ToLower(strings.TrimSpace(file.GetConf().Provider))
	model := strings.ToLower(strings.TrimSpace(file.GetConf().Model))
	providerDriver, err := driver.GetDriver(provider, model, cfg.GetCfg().GetToken(provider))
	if err != nil {
		panic(err)
	}

	var items []models.PromptItem
	for _, c := range compiled.Prompts {
		items = append(items, convCompiledToSend(c))
	}

	toSend := models.PromptToSend{
		Items: items,
		Conf:  file.GetConf(),
		Extra: nil,
	}

	printSep()
	shared.RunPrompt(providerDriver, toSend, shared.DefaultResponseBefore)
}

func printSep() {
	fmt.Println("================")
}

func printInfo(fileInfo prompt.FileInfo) {
	var builder strings.Builder
	if fileInfo.Project != "" {
		builder.WriteString("Project: ")
		builder.WriteString(fileInfo.Project)
		builder.WriteString("\n")
	}
	if fileInfo.Version != "" {
		builder.WriteString("Version: ")
		builder.WriteString(fileInfo.Version)
		builder.WriteString("\n")
	}
	if fileInfo.Author != "" {
		builder.WriteString("Author: ")
		builder.WriteString(fileInfo.Author)
		builder.WriteString("\n")
	}
	if fileInfo.License != "" {
		builder.WriteString("License: ")
		builder.WriteString(fileInfo.License)
		builder.WriteString("\n")
	}
	if builder.Len() > 0 {
		fmt.Print(builder.String())
	} else {
		fmt.Println("No info provided by prompt file")
	}
}

func fetchFile(file string) (txt string, structured bool, err error) {
	structured = true
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
	structured = false
	txt, err = iox.ReadAllText(file + ".prompt")
	return
}
