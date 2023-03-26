package gpt3_driver

import (
	"github.com/sashabaranov/go-openai"
)

type GPT3Driver struct {
	Client *openai.Client
}

func New(token string) *GPT3Driver {
	return &GPT3Driver{Client: openai.NewClient(token)}
}

func factoryRequest(model string) openai.CompletionRequest {
	req := openai.CompletionRequest{
		Model: model,
	}
	return req
}
