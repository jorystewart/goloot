package helperFunctions

import (
	"os"
	"path/filepath"
)

func GetModuleRoot(dir string) (rootPath string) {
	if dir == "" {
		panic("dir cannot be an empty string")
	}
	dir = filepath.Clean(dir)
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
	}
	return ""
}
