package utils

import "fmt"

func GetUserNameComplement() []string {
	birthdates := GetBirthdays()

	people := []string{}

	for _, birthdate := range birthdates {
		for _, person := range birthdate.People {
			people = append(people, fmt.Sprintf("%s - %s", person.Name, person.Complement))
		}
	}

	return people
}
