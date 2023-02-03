package files

import (
	"os"
	"path/filepath"
)

func GetFilePaths(root string) ([]string, error) {
	var files []string
	
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() { 
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
