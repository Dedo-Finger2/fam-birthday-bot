package utils_test

import (
	"fmt"
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetEnvVariableBotApiToken(t *testing.T) {
	const key = "BOT_API_TOKEN"
	_, err := utils.GetEnvVariable(key)

	expected := err == nil

	if !expected {
		t.Errorf("key '%s' not found.", key)
	}
}

func TestGetEnvVariableAllowedUserIDS(t *testing.T) {
	const key = "ALLOWED_CHAT_IDS"
	_, err := utils.GetEnvVariable(key)

	expected := err == nil

	if !expected {
		t.Errorf("key '%s' not found.", key)
	}
}

func TestGetNonExistingEnvVariable(t *testing.T) {
	const key = "NON_EXISTING"
	_, err := utils.GetEnvVariable(key)

	expected := err.Error() == fmt.Sprintf("key %s not found in .env file", key)

	if !expected {
		t.Error("it was expecting an error, but got nil")
	}
}
