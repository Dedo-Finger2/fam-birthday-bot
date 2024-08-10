package utils

import (
	"strings"
	"time"
)

func GetCurrentDateDDMM() string {
	currentDate := strings.Split(strings.Split(time.Now().Local().String(), " ")[0], "-")[1:]
	formattedDate := currentDate[1] + "-" + currentDate[0]
	return formattedDate
}
