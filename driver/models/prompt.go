package models

type PromptToSend struct {
	Model string
	Items []PromptItem
	Extra any
}

type PromptItem struct {
	Content string
	Role    string
	Extra   any
}
