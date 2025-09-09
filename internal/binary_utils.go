package embed_env_internal

import (
	"os"
	"path"
	"strings"
)

func Folder() string {
	return path.Dir(Uri())
}

func Uri() string {
	uri, _ := os.Executable()
	return uri
}

func Name() string {
	name, _ := strings.CutPrefix(Uri(), Folder())
	return name
}
