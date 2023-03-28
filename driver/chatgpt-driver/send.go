package chatgpt_driver

import (
	"context"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

func (c *ChatGPTDriver) SendRequest(p models.PromptToSend) (*openai.ChatCompletionResponse, error) {
	req := factoryRequest(p)
	if len(req.Messages) == 0 {
		return nil, interfaces.ErrEmptyPrompt
	}
	ctx := context.Background()
	resp, err := c.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ChatGPTDriver) SendStreamRequest(p models.PromptToSend) (*openai.ChatCompletionStream, error) {
	req := factoryRequest(p)
	if len(req.Messages) == 0 {
		return nil, interfaces.ErrEmptyPrompt
	}
	ctx := context.Background()
	return c.Client.CreateChatCompletionStream(ctx, req)
}
