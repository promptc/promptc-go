package gpt3_driver

import (
	"context"
	"errors"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

var EmptyPrompt = errors.New("empty prompt")

func (c *GPT3Driver) SendRequest(p models.PromptToSend) (*openai.CompletionResponse, error) {
	req := factoryRequest(p.Model)
	for _, _p := range p.Items {
		if _p.Content == "" {
			continue
		}
		req.Prompt = _p.Content
	}
	if req.Prompt == "" {
		return nil, EmptyPrompt
	}
	ctx := context.Background()
	resp, err := c.Client.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
