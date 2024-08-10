package utils

import (
	"encoding/json"
	"os"
	"path"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
)

func GetBirthdays() []types.Birthday {
	var birhtdaysPaylod struct {
		Dates []types.Birthday `json:"dates"`
	}

	birthdaysFile, err := os.ReadFile(path.Join(GetCurrentDir(), "internal", "config", "birthdays.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(birthdaysFile, &birhtdaysPaylod)
	if err != nil {
		panic(err)
	}

	return birhtdaysPaylod.Dates
}
