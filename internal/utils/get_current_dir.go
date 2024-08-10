package utils

import "os"

func GetCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentDir
}
