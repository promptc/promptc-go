package shared

import (
	"os"
	"path"
)

func GetUserFolder() string {
	cfgP, _ := os.UserConfigDir()
	return path.Join(cfgP, "promptc", "cli")
}

func InitPath() {
	err := os.MkdirAll(GetUserFolder(), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func GetPath(file string) string {
	return path.Join(GetUserFolder(), file)
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
