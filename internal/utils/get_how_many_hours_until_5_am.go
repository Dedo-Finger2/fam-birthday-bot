package utils

import (
	"math"
	"time"
)

type getHowManyHoursUntil5AMOutput struct {
	Hour         float64
	Minutes      int
	CompleteTime time.Duration
}

func GetHowManyHoursUntil5AM(currentDate time.Time) getHowManyHoursUntil5AMOutput {
	var (
		nextDayHour    = 4
		nextDayMinutes = 58
		nextDayMili    = 0
		nextDayNano    = 0
	)
	nextDay := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day()+1, nextDayHour, nextDayMinutes, nextDayMili, nextDayNano, currentDate.Location())

	timeUntil5AM := nextDay.Sub(currentDate)

	output := getHowManyHoursUntil5AMOutput{
		Hour:         math.Floor(timeUntil5AM.Hours()),
		Minutes:      int(timeUntil5AM.Minutes()) % 60,
		CompleteTime: timeUntil5AM,
	}

	return output
}
