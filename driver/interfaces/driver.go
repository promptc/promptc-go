package interfaces

import (
	"io"
)

type ProviderDriver interface {
	GetResponse(prompt models.PromptToSend) ([]string, error)
}

type ProviderStreamDriver interface {
	GetStreamResponse(prompt models.PromptToSend) (io.ReadCloser, error)
}
