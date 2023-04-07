package provider

type Privider interface {
	GetPrompt(name string) (string, error)
}
