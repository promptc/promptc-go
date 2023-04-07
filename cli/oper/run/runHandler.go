package run

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
	ptProvider "github.com/promptc/promptc-go/prompt/provider"
	"path/filepath"
	"strings"
)

func RunHandler(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: promptc-cli run [prompt-file] [var-file]")
		return
	}
	promptPath := args[0]
	varPath := args[1]
	promptStr, err := iox.ReadAllText(promptPath)
	if err != nil {
		panic(err)
	}
	varStr, err := iox.ReadAllText(varPath)
	if err != nil {
		panic(err)
	}

	varMap := shared.IniToMap(varStr)
	file := prompt.ParseFile(promptStr)
	file.RefProvider = &ptProvider.FileProvider{
		BasePath: filepath.Dir(promptPath),
	}
	provider := strings.ToLower(strings.TrimSpace(file.GetConf().Provider))
	model := strings.ToLower(strings.TrimSpace(file.GetConf().Model))
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
		Conf:  file.GetConf(),
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
