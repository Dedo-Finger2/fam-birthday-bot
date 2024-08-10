package main

import (
	"log/slog"
	"time"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func main() {
	bot := config.SetupBot()

	for {
		now := time.Now().Local()

		if now.Hour() == 19 && now.Minute() == 15 {
			birthdates := utils.GetBirthdays()
			currentDateDDMM := utils.GetCurrentDateDDMM()

			for _, birthdate := range birthdates {
				if birthdate.Date == currentDateDDMM {
					utils.SendMessage(birthdate.Date, bot)
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
