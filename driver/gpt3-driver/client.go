package gpt3_driver

import (
	gpt_provider "github.com/promptc/promptc-go/driver/gpt-provider"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

type GPT3Driver struct {
	Client *openai.Client
}

func NewWithProvider(token, provider string) *GPT3Driver {
	return &GPT3Driver{Client: openai.NewClientWithConfig(gpt_provider.GetProviderConfig(token, provider))}
}

func New(token string) *GPT3Driver {
	return NewWithProvider(token, "openai")
}

func factoryRequest(p models.PromptToSend) openai.CompletionRequest {
	req := openai.CompletionRequest{
		Model: p.Conf.Model,
	}
	if p.Conf.Temperature != nil {
		req.Temperature = *p.Conf.Temperature
	}
	if len(p.Conf.Stop) > 0 {
		req.Stop = p.Conf.Stop
	}
	for _, _p := range p.Items {
		if _p.Content == "" {
			continue
		}
		req.Prompt = _p.Content
	}
	return req
}
