package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetAllowedUserIDS() ([]int64, error) {
	allowedUsersPayload := []int64{}

	allowedIDSFile, err := GetEnvVariable("ALLOWED_CHAT_IDS")
	if err != nil {
		return []int64{}, fmt.Errorf("failed on trying to get .env variable %w", err)
	}

	IDSString := strings.Split(allowedIDSFile, "|")

	for _, id := range IDSString {
		convertedID, err := strconv.Atoi(id)
		if err != nil {
			return []int64{}, fmt.Errorf("failed to convert id '%s' to int64", id)
		}

		allowedUsersPayload = append(allowedUsersPayload, int64(convertedID))
	}

	return allowedUsersPayload, nil
}
