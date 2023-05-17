package gpt_provider

import "github.com/sashabaranov/go-openai"

// GetProviderConfig returns the openai.ClientConfig based on the token and provider arguments.
//
// token: string representing the token to be used.
// provider: string representing the provider to use, either "openai" or "openai-sb".
// Returns openai.ClientConfig object.
func GetProviderConfig(token string, provider string) openai.ClientConfig {
	conf := openai.DefaultConfig(token)
	if provider == "openai-sb" {
		conf.BaseURL = "https://api.openai-sb.com/v1"
	}
	return conf
}
