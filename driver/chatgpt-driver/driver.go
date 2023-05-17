package chatgpt_driver

import (
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
)

// GetResponse returns a slice of strings and an error. It takes in a prompt of type models.PromptToSend
// and sends the prompt as a request to the ChatGPTDriver. It then retrieves the response and extracts
// the choices from the response to return as a slice of strings. Returns nil and an error if there's an issue
// sending the request or extracting the choices.
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
	return true
}

func (c *ChatGPTDriver) ToStream() interfaces.ProviderStreamDriver {
	return c
}

var _ interfaces.ProviderDriver = (*ChatGPTDriver)(nil)
