package utils

import (
	"fmt"
	"log/slog"
	"math"
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

		currentDateDDMM := GetCurrentDateMMDD()

		for _, birthDate := range birthDates {
			if birthDate.Date == currentDateDDMM {
				SendMessage(birthDate.Date, bot)
				isBirthday = true
				break
			}
		}

		if isBirthday {
			slog.Info("Messages sent! Next validation in 24 hours.")
		} else {
			slog.Info("There is not birthday today... Next validation in 24 hours.")
		}

		time.Sleep(time.Hour * 24)
	case hour > 5:
		birthDates, err := GetBirthDatesJson()
		if err != nil {
			slog.Error("Error trying to get birthdays.", "error", err)
			os.Exit(1)
		}

		currentDateDDMM := GetCurrentDateMMDD()

		for _, birthDate := range birthDates {
			if birthDate.Date == currentDateDDMM {
				SendMessage(birthDate.Date, bot)
				isBirthday = true
				break
			}
		}

		currentDate := time.Now().Local()
		var (
			nextDayHour    = 4
			nextDayMinutes = 58
			nextDayMili    = 0
			nextDayNano    = 0
		)
		nextDay := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day()+1, nextDayHour, nextDayMinutes, nextDayMili, nextDayNano, currentDate.Location())

		timeUntilNextValidation := nextDay.Sub(currentDate)

		if isBirthday {
			slog.Info(fmt.Sprintf("Messages sent! Next validation in %0.f hours and %0.f minutes.", math.Floor(timeUntilNextValidation.Hours()), math.Floor(timeUntilNextValidation.Minutes()/60)))
		} else {
			slog.Info(fmt.Sprintf("There is not birthday today... Next validation in %0.f hours and %0.f minutes.", math.Floor(timeUntilNextValidation.Hours()), math.Floor(timeUntilNextValidation.Minutes()/60)))
		}
		time.Sleep(timeUntilNextValidation)
	default:
		slog.Warn("It is not 5 am yet, waiting 1 minute before trying again...")
		time.Sleep(time.Minute)
	}
}
