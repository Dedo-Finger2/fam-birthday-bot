package main

import (
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/config"
	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func main() {
	bot := config.SetupBot()

	birthdates := utils.GetBirthdays()

	currentDateDDMM := utils.GetCurrentDateDDMM()

	for _, birthdate := range birthdates {
		if birthdate.Date == currentDateDDMM {
			utils.SendMessage(birthdate.Date, bot)
			return
		}
	}
}
