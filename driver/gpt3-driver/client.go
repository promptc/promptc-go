package gpt3_driver

import (
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

type GPT3Driver struct {
	Client *openai.Client
}

func New(token string) *GPT3Driver {
	return &GPT3Driver{Client: openai.NewClient(token)}
}

func factoryRequest(p models.PromptToSend) openai.CompletionRequest {
	req := openai.CompletionRequest{
		Model: p.Model,
	}
	for _, _p := range p.Items {
		if _p.Content == "" {
			continue
		}
		req.Prompt = _p.Content
	}
	return req
}
