package run

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/chzyer/readline"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
	ptProvider "github.com/promptc/promptc-go/prompt/provider"
	"github.com/promptc/promptc-go/utils"
	"path/filepath"
	"strings"
)

func SimpleRunHandler(args []string) {
	if len(args) < 1 {
		panic("Usage: promptc-cli [prompt-file] [input?]")
	}
	path := args[0]
	input := ""
	inputs := args[1:]
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
		fmt.Println("Please enter following vars:")
		for k, v := range file.VarConstraint {
			input, err := readline.Line(k + " (" + v.Type() + "): ")
			if err != nil {
				panic(err)
			}
			varMap[k] = input
		}
		//panic("Too many vars")
	} else if len(file.Vars) == 1 {
		for k, _ := range file.Vars {
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
	for k, v := range varMap {
		fmt.Println("VAR.", k, " -> ", v)
	}
	compiled := file.Compile(varMap)
	fmt.Println("Compiled To: ")
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

func printInfo(f prompt.FileInfo) {
	sb := strings.Builder{}
	if f.Project != "" {
		sb.WriteString("Project: ")
		sb.WriteString(f.Project)
		sb.WriteString("\n")
	}
	if f.Version != "" {
		sb.WriteString("Version: ")
		sb.WriteString(f.Version)
		sb.WriteString("\n")
	}
	if f.Author != "" {
		sb.WriteString("Author: ")
		sb.WriteString(f.Author)
		sb.WriteString("\n")
	}
	if f.License != "" {
		sb.WriteString("License: ")
		sb.WriteString(f.License)
		sb.WriteString("\n")
	}
	if sb.Len() > 0 {
		fmt.Print(sb.String())
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
