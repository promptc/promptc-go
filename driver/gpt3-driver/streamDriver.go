package gpt3_driver

import (
	"errors"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/sashabaranov/go-openai"
	"io"
)

type GPT3Receiver struct {
	Stream *openai.CompletionStream
}

func (r *GPT3Receiver) Receive() (choices []string, err error, eof bool) {
	resp, err := r.Stream.Recv()
	if errors.Is(err, io.EOF) {
		return nil, err, true
	}

	if err != nil {
		return nil, err, false
	}
	return choicesToArr(resp.Choices), nil, false
}

func (r *GPT3Receiver) Close() {
	r.Stream.Close()
}

var _ interfaces.ProviderStreamDriver = (*GPT3Driver)(nil)

func (c *GPT3Driver) GetStreamResponse(prompt models.PromptToSend) (interfaces.StreamReceiver, error) {
	resp, err := c.SendStreamRequest(prompt)
	if err != nil {
		return nil, err
	}
	return &GPT3Receiver{Stream: resp}, nil
}
