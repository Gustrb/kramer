package utils

import (
	"os"
	"path/filepath"
)

func GetFolderSize(folderPath string) (uint64, error) {
	var size uint64

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += uint64(info.Size())
		}

		return nil
	})

	return size, err
}
