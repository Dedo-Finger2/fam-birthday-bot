package utils

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(date string, bot *tgbotapi.BotAPI) error {
	allowedUserIDS, err := GetAllowedUserIDS()
	if err != nil {
		return err
	}

	currentDate := time.Now().Local().Format(time.DateOnly)

	userNameComplements, err := GetUserNameComplement(date)
	if err != nil {
		return err
	}

	messageText := fmt.Sprintf("*🎊 Hoje, %s, temos aniversariantes! 🎊* \n\n*Não esqueça* de desejar feliz aniversário para: \n\n*%s*",
		currentDate,
		strings.Join(userNameComplements, "\n"),
	)

	for _, userID := range allowedUserIDS {
		message := tgbotapi.NewMessage(userID, messageText)
		message.AllowSendingWithoutReply = true
		message.ParseMode = "markdown"

		bot.Send(message)
	}

	return nil
}
