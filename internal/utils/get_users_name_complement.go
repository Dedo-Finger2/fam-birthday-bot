package utils

import (
	"fmt"
)

func GetUserNameComplement(date string) ([]string, error) {
	birthDates, err := GetBirthDatesJson()
	if err != nil {
		return []string{}, err
	}

	people := []string{}

	for _, birthDate := range birthDates {
		if birthDate.Date != date {
			continue
		}

		for _, person := range birthDate.People {
			people = append(people, fmt.Sprintf("%s - %s", person.Name, person.Complement))
		}
	}

	return people, nil
}
