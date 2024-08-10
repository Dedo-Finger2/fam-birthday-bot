package utils

import "fmt"

func GetUserNameComplement(date string) []string {
	birthdates := GetBirthdays()

	people := []string{}

	for _, birthdate := range birthdates {
		if birthdate.Date != date {
			continue
		}

		for _, person := range birthdate.People {
			people = append(people, fmt.Sprintf("%s - %s", person.Name, person.Complement))
		}
	}

	return people
}
