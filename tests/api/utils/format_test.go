package utils

import (
	"mercafacil-challenge/api/utils"
	"testing"
)

func TestFormatErrorCellphone(t *testing.T) {
	expected := "Cellphone invalid"
	result := utils.FormatError("Error reading cellphone field")
	if result.Error() != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestFormatErrorAnotherReason(t *testing.T) {
	expected := "An error has occurred"
	result := utils.FormatError("Name invalid")
	if result.Error() != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}
