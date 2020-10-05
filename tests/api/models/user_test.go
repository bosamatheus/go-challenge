package models

import (
	"mercafacil-challenge/api/models"
	"testing"
)

func TestFormatCellphone(t *testing.T) {
	expected := "+55 (41) 91234-5678"
	result := models.FormatCellphone("5541912345678")
	if result != expected {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}
