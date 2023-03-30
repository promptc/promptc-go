package run

import (
	"fmt"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
	"io"
	"os"
	"strings"
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
	file := prompt.ParseFile(string(promptBs))
	provider := strings.ToLower(strings.TrimSpace(file.Conf.Provider))
	model := strings.ToLower(strings.TrimSpace(file.Conf.Model))
	providerDriver, err := driver.GetDriver(provider, model, cfg.GetCfg().GetToken(provider))
	if err != nil {
		panic(err)
	}
	compiled := file.Compile(varMap)
	var items []models.PromptItem
	for _, c := range compiled.Prompts {
		items = append(items, convCompiledToSend(c))
	}
	toSend := models.PromptToSend{
		Items: items,
		Model: model,
		Extra: nil,
	}
	shared.RunPrompt(providerDriver, toSend, shared.DefaultResponseBefore)
}

func convCompiledToSend(c prompt.CompiledPrompt) models.PromptItem {
	return models.PromptItem{
		Content: c.Prompt,
		Extra:   c.Extra,
	}
}
