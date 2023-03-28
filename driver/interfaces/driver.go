package interfaces

import (
	"github.com/promptc/promptc-go/driver/models"
)

type ProviderDriver interface {
	GetResponse(prompt models.PromptToSend) ([]string, error)
	StreamAvailable() bool
	ToStream() ProviderStreamDriver
}

type ProviderStreamDriver interface {
	GetStreamResponse(prompt models.PromptToSend) (StreamReceiver, error)
}

type StreamReceiver interface {
	Receive() (choices []string, err error, eof bool)
	Close()
}
