package cfg

import (
	"fmt"

	"github.com/hjson/hjson-go/v4"
)

// SetHandler sets a key and value in the configuration file.
//
// It takes in an array of strings containing the key and value to be set.
// This function does not return anything.
func SetHandler(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: promptc-cli set [key] [value]")
		return
	}
	cfg := GetCfg()
	cfgM := make(map[string]string)
	bs, _ := hjson.Marshal(*cfg)
	_ = hjson.Unmarshal(bs, &cfgM)
	if len(cfgM) == 0 {
		cfgM = make(map[string]string)
	}
	cfgM[args[0]] = args[1]
	bs, _ = hjson.Marshal(cfgM)
	_ = hjson.Unmarshal(bs, cfg)
	Save()
}

func ShowHandler(args []string) {
	cfg := GetCfg()
	bs, _ := hjson.Marshal(*cfg)
	fmt.Println(string(bs))
}
