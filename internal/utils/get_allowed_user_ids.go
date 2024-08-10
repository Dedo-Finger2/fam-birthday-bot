package utils

import (
	"encoding/json"
	"os"
	"path"
)

func GetAllowedUserIDS() []int64 {
	var allowedUsersPayload struct {
		IDS []int64 `json:"allowed_ids"`
	}

	allowedIDSFile, err := os.ReadFile(path.Join(GetCurrentDir(), "internal", "config", "allowed_user_ids.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(allowedIDSFile, &allowedUsersPayload)
	if err != nil {
		panic(err)
	}

	return allowedUsersPayload.IDS
}
