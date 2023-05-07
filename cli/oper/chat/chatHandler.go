package chat

import (
	"errors"
	"fmt"
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/chzyer/readline"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
)

func ChatHandler(args []string) {
	var userInput []string
	var gptInput []string

	userInputHint := "YOU> "

	colourPrefix, needReset := console.Green.AsForeground().ConsoleString()
	userInputHint = colourPrefix + userInputHint
	if needReset {
		userInputHint += console.ResetColourSymbol
	}

	rl, err := readline.New(userInputHint)
	if err != nil {
		panic(err)
	}
	providerDriver := driver.GetOpenAIDriver(cfg.GetCfg().DefaultProvider, cfg.GetCfg().OpenAIToken)
	for {
		line, err := rl.Readline()
		if err != nil {
			if errors.Is(err, readline.ErrInterrupt) {
				break
			}
			fmt.Println(err)
			break
		}
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
					"role": "assistant",
				},
			})
		}
	}
	toSend := models.PromptToSend{
		Items: prmpt,
		Conf: prompt.Conf{
			Model: "gpt-3.5-turbo",
		},
	}
	rst := shared.RunPrompt(drv, toSend, func(id int) {
		console.Blue.AsForeground().Write("GPT> ")
	})
	gptInput = append(gptInput, rst[0])
	return gptInput
}
