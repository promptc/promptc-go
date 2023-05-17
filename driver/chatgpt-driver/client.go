package chatgpt_driver

import (
	gpt_provider "github.com/promptc/promptc-go/driver/gpt-provider"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

type ChatGPTDriver struct {
	Client *openai.Client
}

func NewWithProvider(token, provider string) *ChatGPTDriver {
	return &ChatGPTDriver{
		Client: openai.NewClientWithConfig(gpt_provider.GetProviderConfig(token, provider)),
	}
}

func New(token string) *ChatGPTDriver {
	return NewWithProvider(token, "openai")
}

// factoryRequest returns an openai.ChatCompletionRequest object based on the provided prompt.
//
// Takes in a models.PromptToSend object as a parameter and returns an openai.ChatCompletionRequest object.
func factoryRequest(p models.PromptToSend) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model:    p.Conf.Model,
		Messages: getMessages(p),
	}
	if p.Conf.Temperature != nil {
		req.Temperature = *p.Conf.Temperature
	}
	if len(p.Conf.Stop) > 0 {
		req.Stop = p.Conf.Stop
	}
	return req
}

// getMessages returns a list of chat messages to send using OpenAI API.
//
// It takes a `models.PromptToSend` struct as its only parameter, with the prompt
// items (strings) and their respective `extra` keys to determine their role.
// The function returns a slice of `openai.ChatCompletionMessage` structs
// containing the role and content of each prompt.
func getMessages(p models.PromptToSend) []openai.ChatCompletionMessage {
	var messages []openai.ChatCompletionMessage
	for _, item := range p.Items {
		if item.Content == "" {
			continue
		}

		role := "user"
		if extra, ok := item.Extra["role"].(string); ok {
			role = extra
		}

		msg := openai.ChatCompletionMessage{
			Role:    role,
			Content: item.Content,
		}
		messages = append(messages, msg)
	}
	return messages
}
