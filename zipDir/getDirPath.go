package zipdir

import (
	"os"
	"path/filepath"
)

func GetDirPath(root string) (files map[string]string) {
	files = make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != root {
			files[info.Name()] = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}
