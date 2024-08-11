package utils

import (
	"fmt"
	"os"
)

func GetCurrentDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Failed to get current directory. %w", err)
	}
	return currentDir, nil
}
