package interfaces

import (
	"github.com/promptc/promptc-go/driver/models"
	"io"
)

type ProviderDriver interface {
	GetResponse(prompt models.PromptToSend) ([]string, error)
}

type ProviderStreamDriver interface {
	GetStreamResponse(prompt models.PromptToSend) (io.ReadCloser, error)
}
