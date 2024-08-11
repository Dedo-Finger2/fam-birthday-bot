package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) (string, error) {
	absPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	configFilePath := path.Join(absPath, ".env")
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("failed to read .env file: %w", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		return "", fmt.Errorf("key %s not found in .env file", key)
	}
	if value == "" {
		return "", fmt.Errorf("key %s has no value defined", key)
	}

	return value, nil
}
