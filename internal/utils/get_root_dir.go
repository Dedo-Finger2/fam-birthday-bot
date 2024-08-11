package utils

import (
	"os"
	"path/filepath"
)

func GetRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)

		if parentDir == dir {
			break
		}

		dir = parentDir
	}

	return "", os.ErrNotExist
}
