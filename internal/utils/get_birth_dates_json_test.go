package utils_test

import (
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetBirthDatesJson(t *testing.T) {
	output, err := utils.GetBirthDatesJson()
	if err != nil {
		t.Errorf("failed to get birth dates: %s", err)
	}

	expect := len(output) > 0

	if !expect {
		t.Errorf("invalid length, got 0")
	}
}
