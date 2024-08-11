package utils_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetCurrentDateMMDD(t *testing.T) {
	var currentDate = time.Now().Local().Format(time.DateOnly)
	output := utils.GetCurrentDateMMDD()

	expected := strings.Contains(currentDate, output)

	if !expected {
		t.Errorf("invalid date returned, today is %s", currentDate)
	}
}
