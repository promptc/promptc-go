package gpt_provider

import "github.com/sashabaranov/go-openai"

func GetProviderConfig(token string, provider string) openai.ClientConfig {
	conf := openai.DefaultConfig(token)
	if provider == "openai" {
		return conf
	}
	switch provider {
	case "openai-sb":
		conf.BaseURL = "https://api.openai-sb.com/v1"
	}
	return conf
}
