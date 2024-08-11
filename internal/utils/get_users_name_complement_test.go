package utils_test

import (
	"strings"
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetUserNameComplement(t *testing.T) {
	const date = "99-99" // Date exclusive for testing purpose
	output, err := utils.GetUserNameComplement(date)
	if err != nil {
		t.Errorf("failed to get output: %s", err)
	}

	isOutputLengthHigherThanZero := len(output) > 0

	expected := isOutputLengthHigherThanZero

	if !expected {
		t.Error("no data was returned")
	}
}

func TestGetUserNameComplementInvalidDate(t *testing.T) {
	const date = "10000-10000"
	output, err := utils.GetUserNameComplement(date)
	if err != nil {
		t.Errorf("failed to get output: %s", err)
	}

	expected := len(output) == 0

	if !expected {
		t.Error("expected to return an empty slice")
	}
}

func TestGetUserNameComplementFormat(t *testing.T) {
	const date = "99-99" // Date exclusive for testing purpose
	output, err := utils.GetUserNameComplement(date)
	if err != nil {
		t.Errorf("failed to get output: %s", err)
	}

	expected := len(strings.Split(output[0], " - ")) == 2

	if !expected {
		t.Error("expected to return a string with '-' as a separator")
	}
}

func TestGetUserNameComplementFormatWithMultipleValues(t *testing.T) {
	const date = "99-99" // Date exclusive for testing purpose
	output, err := utils.GetUserNameComplement(date)
	if err != nil {
		t.Errorf("failed to get output: %s", err)
	}

	areAllValuesValid := []bool{}

	for _, value := range output {
		if len(strings.Split(value, " - ")) == 2 {
			areAllValuesValid = append(areAllValuesValid, true)
		} else {
			areAllValuesValid = append(areAllValuesValid, false)
		}
	}

	for _, isValueValid := range areAllValuesValid {
		if !isValueValid {
			t.Error("not all values meet the wanted format")
		}
	}
}
