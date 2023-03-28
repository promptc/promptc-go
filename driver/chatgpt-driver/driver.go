package chatgpt_driver

import (
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
)

func (c *ChatGPTDriver) GetResponse(prompt models.PromptToSend) ([]string, error) {
	resp, err := c.SendRequest(prompt)
	if err != nil {
		return nil, err
	}
	var choices []string
	for _, choice := range resp.Choices {
		choices = append(choices, choice.Message.Content)
	}
	return choices, nil
}

func (c *ChatGPTDriver) StreamAvailable() bool {
	return false
}

func (c *ChatGPTDriver) ToStream() interfaces.ProviderStreamDriver {
	return c
}

var _ interfaces.ProviderDriver = (*ChatGPTDriver)(nil)
