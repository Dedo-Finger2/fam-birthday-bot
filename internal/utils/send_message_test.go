package utils_test

import (
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestSendMessage(t *testing.T) {
	userIDS, err := utils.GetTestingAllowedUserIDS()
	if err != nil {
		t.Errorf("Failed to get user allowed ids: %s", err.Error())
	}
	bot, err := config.SetupBot()
	if err != nil {
		t.Errorf("Failed to setup bot: %s", err.Error())
	}

	if err = utils.SendMessage("99-99", bot, userIDS); err != nil {
		t.Errorf("Failed to send messages: %s", err.Error())
	}
}
