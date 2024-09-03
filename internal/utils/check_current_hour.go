package utils

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CheckCurrentHourCron(currentDate time.Time, bot *tgbotapi.BotAPI) {
	isBirthday := false

	birthDates, err := GetBirthDatesYaml()
	if err != nil {
		slog.Error("Error trying to get birthdays.", "error", err)
		os.Exit(1)
	}

	HandleSendingMessageToUsers(birthDates, bot, &isBirthday)

	timeUntilNextValidation := GetHowManyHoursUntil5AM(currentDate)

	if isBirthday {
		slog.Info(fmt.Sprintf("Messages sent! Next validation in %0.f hours and %d minutes.", timeUntilNextValidation.Hour, timeUntilNextValidation.Minutes))
	} else {
		slog.Info(fmt.Sprintf("There is not birthday today... Next validation in %0.f hours and %d minutes.", timeUntilNextValidation.Hour, timeUntilNextValidation.Minutes))
	}
}
