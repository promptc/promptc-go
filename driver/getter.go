package driver

import (
	"errors"
	"github.com/promptc/promptc-go/driver/chatgpt-driver"
	"github.com/promptc/promptc-go/driver/gpt3-driver"
	"github.com/promptc/promptc-go/driver/interfaces"
)

var ErrProviderNotFound = errors.New("provider not found")

func GetDriver(provider, model, token string) (interfaces.ProviderDriver, error) {
	// see https://platform.openai.com/docs/models/model-endpoint-compatibility
	if provider == "openai" {
		switch model {
		case "gpt-4", "gpt-4-0314", "gpt-4-32k", "gpt-4-32k-0314", "gpt-3.5-turbo", "gpt-3.5-turbo-0301":
			return chatgpt_driver.New(token), nil
		case "text-davinci-003", "text-davinci-002", "text-curie-001", "text-babbage-001", "text-ada-001", "davinci", "curie", "babbage", "ada":
			return gpt3_driver.New(token), nil
		}
	}
	return nil, ErrProviderNotFound
}

func GetDefaultDriver(token string) interfaces.ProviderDriver {
	return chatgpt_driver.New(token)
}
