package utils

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "cellphone") {
		return errors.New("Cellphone invalid")
	}
	return errors.New("An error has occurred")
}
