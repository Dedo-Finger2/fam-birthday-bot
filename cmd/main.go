package main

import (
	"log/slog"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
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
	c := cron.New()

	_, err := c.AddFunc("0 5 * * *", func() {
		slog.Info("Running scheduled job at 5:00 AM")
		currentTime := time.Now().Local()
		utils.CheckCurrentHourCron(currentTime, bot)
	})
	if err != nil {
		slog.Error("Failed to schedule cron job", "error", err)
		return
	}

	c.Start()

	select {}
}
