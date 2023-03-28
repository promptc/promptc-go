package gpt3_driver

import (
	"context"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

func (c *GPT3Driver) SendRequest(p models.PromptToSend) (*openai.CompletionResponse, error) {
	req := factoryRequest(p)
	if req.Prompt == "" {
		return nil, interfaces.ErrEmptyPrompt
	}
	ctx := context.Background()
	resp, err := c.Client.CreateCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *GPT3Driver) SendStreamRequest(p models.PromptToSend) (*openai.CompletionStream, error) {
	req := factoryRequest(p)
	if req.Prompt == "" {
		return nil, interfaces.ErrEmptyPrompt
	}
	ctx := context.Background()
	return c.Client.CreateCompletionStream(ctx, req)
}
