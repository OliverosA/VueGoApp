package files

import (
	"os"
	"path/filepath"
)

func GetFilePaths(root string) ([]string, error) {
	var files []string
	
	// walking or reading all the root (directories and files)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// if the file isn't a directory
		if !info.IsDir() { 
			// append the file to the slice
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
