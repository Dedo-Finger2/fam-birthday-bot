package utils_test

import (
	"testing"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetHowManyHoursUntil5AM(t *testing.T) {
	currentDate := time.Now().Local()
	mockTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 17, 0, 0, 0, currentDate.Location())

	output := utils.GetHowManyHoursUntil5AM(mockTime)

	expected := output.Hour == float64(11) && output.Minutes == 58

	if !expected {
		t.Errorf("Test failed, wrong logic when trying to get how many hours until 5am. Expected '%f hours and %d minutes' but got '%f hours and %d minutes'", float64(11), 58, output.Hour, output.Minutes)
	}
}
