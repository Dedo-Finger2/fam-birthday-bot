package config

import (
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"
)

var PermissionIDS []int64 = []int64{6105709676, 7496811705}

<<<<<<< HEAD
type ConfigJsonPaylod struct {
	Dates []types.Birthday `json:"dates"`
=======
func GetBirthdays() []types.Birthday {
	var birhtdaysPaylod struct {
		Dates []types.Birthday `json:"dates"`
	}

	birthdaysFile, err := os.ReadFile(path.Join(currentDir, "internal", "config", "birthdays.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(birthdaysFile, &birhtdaysPaylod)
	if err != nil {
		panic(err)
	}

	return birhtdaysPaylod.Dates
}

func GetAllowedUserIDS() []int64 {
	var allowedUsersPayload struct {
		IDS []int64 `json:"allowed_ids"`
	}

	allowedIDSFile, err := os.ReadFile(path.Join(currentDir, "internal", "config", "allowed_user_ids.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(allowedIDSFile, &allowedUsersPayload)
	if err != nil {
		panic(err)
	}

	return allowedUsersPayload.IDS
>>>>>>> af698d6 (feat: moves configuration files to internal/config dir + updates config.go path to config files)
}
