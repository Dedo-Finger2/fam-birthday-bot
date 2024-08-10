package utils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {
	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigFile(path.Join(absPath, ".env"))

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		panic("key not found.")
	}
	if value == "" {
		panic("key has no value defined.")
	}

	return value
}
