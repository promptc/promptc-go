package gpt3_driver

import (
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
)

func (c *GPT3Driver) GetResponse(prompt models.PromptToSend) ([]string, error) {
	resp, err := c.SendRequest(prompt)
	if err != nil {
		return nil, err
	}
	var choices []string
	for _, choice := range resp.Choices {
		choices = append(choices, choice.Text)
	}
	return choices, nil
}

var _ interfaces.ProviderDriver = (*GPT3Driver)(nil)
