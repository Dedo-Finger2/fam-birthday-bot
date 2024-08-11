package utils_test

import (
	"os"
	"strings"
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetRootDir(t *testing.T) {
	output, err := utils.GetRootDir()
	if err != nil {
		t.Errorf("test failed: %s", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		t.Error("failed to get current dir")
	}

	expected := strings.Contains(currentDir, output) && strings.HasSuffix(output, "-bot")

	if !expected {
		t.Errorf("wrong root path provided, expected it to end with '-bot' but got %s", output)
	}
}
