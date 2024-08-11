package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
)

func GetBirthdays() ([]types.Birthday, error) {
	var birhtdaysPaylod struct {
		Dates []types.Birthday `json:"dates"`
	}

	currentDir, err := GetCurrentDir()
	if err != nil {
		return []types.Birthday{}, err
	}

	birthdaysFile, err := os.ReadFile(path.Join(currentDir, "internal", "config", "birthdays.json"))
	if err != nil {
		return []types.Birthday{}, fmt.Errorf("Failed to read file birthdays.json. %w", err)
	}

	err = json.Unmarshal(birthdaysFile, &birhtdaysPaylod)
	if err != nil {
		return []types.Birthday{}, fmt.Errorf("Failed to Unmarshal json file. %w", err)
	}

	return birhtdaysPaylod.Dates, nil
}
