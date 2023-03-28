package chatgpt_driver

import (
	"errors"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
	"io"
)

type ChatGPTReceiver struct {
	Stream *openai.ChatCompletionStream
}

// Receive Get Delta from response
func (r *ChatGPTReceiver) Receive() (choices []string, err error, eof bool) {
	resp, err := r.Stream.Recv()
	if errors.Is(err, io.EOF) {
		return nil, err, true
	}

	if err != nil {
		return nil, err, false
	}
	for _, choice := range resp.Choices {
		choices = append(choices, choice.Delta.Content)
	}
	return choices, nil, false
}

func (r *ChatGPTReceiver) Close() {
	r.Stream.Close()
}

var _ interfaces.ProviderStreamDriver = (*ChatGPTDriver)(nil)

func (c *ChatGPTDriver) GetStreamResponse(prompt models.PromptToSend) (interfaces.StreamReceiver, error) {
	resp, err := c.SendStreamRequest(prompt)
	if err != nil {
		return nil, err
	}
	return &ChatGPTReceiver{Stream: resp}, nil
}
