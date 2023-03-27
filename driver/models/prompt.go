package models

type PromptToSend struct {
	Model string
	Items []PromptItem
	Extra map[string]any
}

type PromptItem struct {
	Content string
	Extra   map[string]any
}
