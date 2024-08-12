package main

import (
	"log/slog"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error

	bot, err = config.SetupBot()
	if err != nil {
		slog.Error("Error trying to create a new bot API.", "error", err)
		return
	}
}

func main() {
	// INF Loop
	for {
		currentHour := time.Now().Local().Hour()
		utils.CheckCurrentHourCron(currentHour, bot)
	}
}
