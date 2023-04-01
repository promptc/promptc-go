package models

import "github.com/promptc/promptc-go/prompt"

type PromptToSend struct {
	Conf  prompt.Conf
	Items []PromptItem
	Extra map[string]any
}

type PromptItem struct {
	Content string
	Extra   map[string]any
}
