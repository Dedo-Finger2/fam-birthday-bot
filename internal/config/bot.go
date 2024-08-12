package config

import (
	"log/slog"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupBot() (*tgbotapi.BotAPI, error) {
	botToken, err := utils.GetEnvVariable("BOT_API_TOKEN")
	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	slog.Info("Connected to bot!", "user-name", bot.Self.UserName)

	return bot, nil
}
