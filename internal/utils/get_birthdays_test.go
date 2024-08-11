package utils_test

import (
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetBirthdays(t *testing.T) {
	output, err := utils.GetBirthdays()
	if err != nil {
		t.Errorf("failed to get birthdays: %s", err)
	}

	expect := len(output) > 0

	if !expect {
		t.Errorf("invalid length, got 0")
	}
}
