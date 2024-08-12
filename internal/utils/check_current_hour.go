package utils

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CheckCurrentHourCron(hour int, bot *tgbotapi.BotAPI) {
	isBirthday := false

	switch {
	case hour == 5:
		birthDates, err := GetBirthDatesJson()
		if err != nil {
			slog.Error("Error trying to get birthdays.", "error", err)
			os.Exit(1)
		}

		HandleSendingMessageToUsers(birthDates, bot, &isBirthday)

		if isBirthday {
			slog.Info("Messages sent! Next validation in 24 hours.")
		} else {
			slog.Info("There is not birthday today... Next validation in 24 hours.")
		}

		time.Sleep(time.Hour * 24) // 1 day
	case hour > 5:
		birthDates, err := GetBirthDatesJson()
		if err != nil {
			slog.Error("Error trying to get birthdays.", "error", err)
			os.Exit(1)
		}

		HandleSendingMessageToUsers(birthDates, bot, &isBirthday)

		timeUntilNextValidation := GetHowManyHoursUntil5AM()

		if isBirthday {
			slog.Info(fmt.Sprintf("Messages sent! Next validation in %0.f hours and %0.f minutes.", timeUntilNextValidation.Hour, timeUntilNextValidation.Minutes))
		} else {
			slog.Info(fmt.Sprintf("There is not birthday today... Next validation in %0.f hours and %0.f minutes.", timeUntilNextValidation.Hour, timeUntilNextValidation.Minutes))
		}

		time.Sleep(timeUntilNextValidation.CompleteTime)
	default:
		slog.Warn("It is not 5 am yet, waiting 1 minute before trying again...")
		time.Sleep(time.Minute)
	}
}
