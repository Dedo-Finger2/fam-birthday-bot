package utils_test

import (
	"reflect"
	"testing"

	"github.com/Dedo-Finger2/fam-birthday-bot/internal/utils"
)

func TestGetAllowedUserIDS(t *testing.T) {
	const wantedType = "[]int64"
	output, err := utils.GetAllowedUserIDS()
	if err != nil {
		t.Error()
	}
	outputType := reflect.TypeOf(output).String()

	expectedType := outputType == wantedType
	expectedLength := len(output) > 0

	if !expectedType {
		t.Errorf("invalid type, wanted '%s' but got '%s'", wantedType, outputType)
	}
	if !expectedLength {
		t.Errorf("invalid length, got 0")
	}
}
