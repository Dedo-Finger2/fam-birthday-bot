package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
)

func GetBirthDatesJson() ([]types.BirthDate, error) {
	var birthDatePayload types.BirthDatesPayload

	rootDir, err := GetRootDir()
	if err != nil {
		return []types.BirthDate{}, err
	}

	jsonFile, err := os.ReadFile(path.Join(rootDir, "internal", "config", "birth_dates.json"))
	if err != nil {
		return []types.BirthDate{}, fmt.Errorf("failed to read file birth_dates.json. %w", err)
	}

	err = json.Unmarshal(jsonFile, &birthDatePayload)
	if err != nil {
		return []types.BirthDate{}, fmt.Errorf("failed to Unmarshal json file. %w", err)
	}

	return birthDatePayload.Dates, nil
}
