package chat

import (
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
)

func ChatHandler(args []string) {
	var userInput []string
	var gptInput []string
	providerDriver := driver.GetDefaultDriver(cfg.GetCfg().GetToken("openai"))
	for {
		console.Green.AsForeground().Write("YOU> ")
		line, _ := console.ReadLine()
		userInput = append(userInput, line)
		gptInput = runGPT(providerDriver, userInput, gptInput)
	}
}

func runGPT(drv interfaces.ProviderDriver, userInput []string, gptInput []string) []string {
	var prmpt []models.PromptItem
	for i, line := range userInput {
		prmpt = append(prmpt, models.PromptItem{
			Content: line,
			Extra: map[string]any{
				"role": "user",
			},
		})
		if i < len(gptInput) {
			prmpt = append(prmpt, models.PromptItem{
				Content: gptInput[i],
				Extra: map[string]any{
					"role": "system",
				},
			})
		}
	}
	toSend := models.PromptToSend{
		Items: prmpt,
		Model: "gpt-3.5-turbo",
	}
	rst := shared.RunPrompt(drv, toSend, func(id int) {
		console.Blue.AsForeground().Write("GPT> ")
	})
	gptInput = append(gptInput, rst[0])
	return gptInput
}
