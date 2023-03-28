package gpt3_driver

import (
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

func (c *GPT3Driver) GetResponse(prompt models.PromptToSend) ([]string, error) {
	resp, err := c.SendRequest(prompt)
	if err != nil {
		return nil, err
	}
	return choicesToArr(resp.Choices), nil
}

var _ interfaces.ProviderDriver = (*GPT3Driver)(nil)

func choicesToArr(choices []openai.CompletionChoice) []string {
	var arr []string
	for _, choice := range choices {
		arr = append(arr, choice.Text)
	}
	return arr
}
