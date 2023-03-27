package cfg

import (
	"github.com/hjson/hjson-go/v4"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"io"
	"os"
)

type Model struct {
	OpenAIToken string `json:"openai_token"`
}

var model *Model

func Save() {
	bs, err := hjson.Marshal(model)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(shared.GetPath("config.json"), bs, 0644)
	if err != nil {
		panic(err)
	}
}

func (m *Model) GetToken(name string) string {
	switch name {
	case "openai":
		return m.OpenAIToken
	}
	return ""
}

func GetCfg() *Model {
	path := shared.GetPath("config.json")
	if model == nil {
		var file *os.File
		var err error
		if !shared.FileExists(path) {
			file, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		} else {
			file, err = os.Open(path)
			if err != nil {
				panic(err)
			}
		}
		bs, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			panic(err)
		}
		if len(string(bs)) == 0 {
			model = defaultModel()
			Save()
		} else {
			err = hjson.Unmarshal(bs, &model)
			if err != nil {
				panic(err)
			}
		}

	}
	return model
}

func defaultModel() *Model {
	return &Model{
		OpenAIToken: "",
	}
}
