package provider

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	"path"
)

type FileProvider struct {
	BasePath string
}

func (f *FileProvider) GetPrompt(name string) (string, error) {
	p := path.Join(f.BasePath, name)
	return iox.ReadAllText(p)
}
