package utils

import (
	"os"
	"path/filepath"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
	"gopkg.in/yaml.v3"
)

func GetBirthDatesYaml() ([]types.BirthDate, error) {
	var birthDatePayload types.BirthDatesPayload

	rootDir, err := GetRootDir()
	if err != nil {
		return []types.BirthDate{}, err
	}

	yamlFile, err := os.ReadFile(filepath.Join(rootDir, "internal", "config", "birth_dates.yaml"))
	if err != nil {
		return []types.BirthDate{}, err
	}

	err = yaml.Unmarshal(yamlFile, &birthDatePayload)
	if err != nil {
		return []types.BirthDate{}, err
	}

	return birthDatePayload.Dates, nil
}
