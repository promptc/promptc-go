package run

import (
	"fmt"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"strings"
)

func BlankHandler(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: promptc-cli blank [prompt]")
		return
	}
	providerDriver := driver.GetDefaultDriver(cfg.GetCfg().GetToken("openai"))
	toSend := models.PromptToSend{
		Items: []models.PromptItem{
			{
				Content: strings.Join(args, ""),
			},
		},
		Model: "gpt-3.5-turbo",
		Extra: nil,
	}
	shared.RunPrompt(providerDriver, toSend, shared.DefaultResponseBefore)
}
