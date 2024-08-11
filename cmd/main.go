package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func main() {
	bot, err := config.SetupBot()
	if err != nil {
		slog.Error("Error trying to create a new bot API.", "error", err)
		return
	}

	for {
		now := time.Now().Local()

		if now.Hour() == 5 {
			birthDates, err := utils.GetBirthdays()
			if err != nil {
				slog.Error("Error trying to get birthdays.", "error", err)
				os.Exit(1)
			}

			currentDateDDMM := utils.GetCurrentDateMMDD()

			for _, birthDate := range birthDates {
				if birthDate.Date == currentDateDDMM {
					utils.SendMessage(birthDate.Date, bot)
					break
				}
			}

			slog.Info("Messages sent! Next validation in 24 hours.")
			time.Sleep(time.Hour * 24)
		} else {
			slog.Warn("it is not 5 am yet, waiting 1 minute before trying again...")
			time.Sleep(time.Minute)
		}
	}
}
