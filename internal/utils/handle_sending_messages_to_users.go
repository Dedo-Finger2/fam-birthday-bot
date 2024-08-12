package utils

import (
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleSendingMessageToUsers(birthDates []types.BirthDate, bot *tgbotapi.BotAPI, isBirthday *bool) {
	currentDateDDMM := GetCurrentDateMMDD()

	for _, birthDate := range birthDates {
		if birthDate.Date == currentDateDDMM {
			SendMessage(birthDate.Date, bot)
			*isBirthday = true
			break
		}
	}
}
