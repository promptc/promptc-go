package chatgpt_driver

import (
	"context"
	"errors"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
)

var EmptyPrompt = errors.New("empty prompt")

func (c *ChatGPTDriver) SendRequest(p models.PromptToSend) (*openai.ChatCompletionResponse, error) {
	req := factoryRequest(p.Model)
	var message []openai.ChatCompletionMessage
	for _, _p := range p.Items {
		if _p.Content == "" {
			continue
		}

		role := "user"
		if len(_p.Extra) > 0 {
			ok := false
			var a any
			a, ok = _p.Extra["role"]
			if ok {
				role, ok = a.(string)
				if !ok {
					role = "user"
				}
			}
		}

		content := openai.ChatCompletionMessage{
			Role:    role,
			Content: _p.Content,
		}
		message = append(message, content)
	}
	if message == nil || len(message) == 0 {
		return nil, EmptyPrompt
	}
	req.Messages = message
	ctx := context.Background()
	resp, err := c.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
