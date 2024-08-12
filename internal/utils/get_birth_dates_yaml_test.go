package utils_test

import (
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetBirthDatesYaml(t *testing.T) {
	content, err := utils.GetBirthDatesYaml()
	if err != nil {
		t.Errorf("failed to get birthDates file content: %s", err)
	}

	expected := len(content) > 0 && content[0].Date == "99-99"

	if !expected {
		t.Error("expected to get birthDates but got an empty struct")
	}
}
