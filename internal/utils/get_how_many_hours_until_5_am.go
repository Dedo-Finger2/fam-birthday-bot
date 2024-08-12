package utils

import (
	"math"
	"time"
)

type getHowManyHoursUntil5AMOutput struct {
	Hour         float64
	Minutes      float64
	CompleteTime time.Duration
}

func GetHowManyHoursUntil5AM() getHowManyHoursUntil5AMOutput {
	currentDate := time.Now().Local()
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
		Minutes:      math.Floor(timeUntil5AM.Minutes() / 60),
		CompleteTime: timeUntil5AM,
	}

	return output
}
