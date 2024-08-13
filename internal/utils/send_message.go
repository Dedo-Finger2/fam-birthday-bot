package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(date string, bot *tgbotapi.BotAPI, allowedUserIDS []int64) error {
	currentDate := time.Now().Local().Format(time.DateOnly)
	failedMessage := map[int64]string{}

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

		_, err := bot.Send(message)
		if err != nil {
			slog.Warn("Failed to send message.", "error", err)
			failedMessage[userID] = err.Error()
		}
	}

	if len(failedMessage) > 0 {
		return errors.New("failed to send messages to some users")
	}

	return nil
}
