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

func getMessages(p models.PromptToSend) []openai.ChatCompletionMessage {
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
	return message
}
