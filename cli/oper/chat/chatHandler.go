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

// ChatHandler handles the chat input and generates a response using the OpenAI GPT driver.
//
// args: A slice of strings containing the input arguments for the function.
// No return type.
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

// runGPT generates prompts from user input and gptInput, sends the prompts to the
// GPT driver, and appends the response to gptInput.
//
// driver: the GPT driver to send the prompts to
// userInput: a slice of strings representing user input
// gptInput: a slice of strings representing previous GPT responses
//
// []string: a slice of strings representing the updated GPT response
func runGPT(driver interfaces.ProviderDriver, userInput []string, gptInput []string) []string {
	prompts := generatePrompts(userInput, gptInput)
	promptToSend := models.PromptToSend{
		Items: prompts,
		Conf: prompt.Conf{
			Model: "gpt-3.5-turbo",
		},
	}
	response := shared.RunPrompt(driver, promptToSend, func(id int) {
		console.Blue.AsForeground().Write("GPT> ")
	})
	gptInput = append(gptInput, response[0])
	return gptInput
}

func generatePrompts(userInput []string, gptInput []string) []models.PromptItem {
	var prompts []models.PromptItem
	for i, line := range userInput {
		prompts = append(prompts, models.PromptItem{
			Content: line,
			Extra: map[string]any{
				"role": "user",
			},
		})
		if i < len(gptInput) {
			prompts = append(prompts, models.PromptItem{
				Content: gptInput[i],
				Extra: map[string]any{
					"role": "assistant",
				},
			})
		}
	}
	return prompts
}
