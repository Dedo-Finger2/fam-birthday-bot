package utils

import (
	"strings"
	"time"
)

func GetCurrentDateMMDD() string {
	currentDate := strings.Split(strings.Split(time.Now().Local().String(), " ")[0], "-")[1:]
	formattedDate := strings.Join(currentDate, "-")
	return formattedDate
}
