package chatgpt_driver

import (
	"github.com/sashabaranov/go-openai"
)

type ChatGPTDriver struct {
	Client *openai.Client
}

func New(token string) *ChatGPTDriver {
	return &ChatGPTDriver{Client: openai.NewClient(token)}
}

func factoryRequest(model string) openai.ChatCompletionRequest {
	content := make([]openai.ChatCompletionMessage, 0)
	req := openai.ChatCompletionRequest{
		Model:    model,
		Messages: content,
	}
	return req
}
