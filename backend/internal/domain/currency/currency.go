package currency

import (
	"errors"
	"regexp"
)

type Currency struct {
	Code string `json:"Code"`
}

func New(code string) (*Currency, error) {
	if len(code) != 3 {
		return nil, errors.New("currency must be 3 characters long")
	}

	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !re.MatchString(code) {
		return nil, errors.New("currency Code must only contain alphabetic characters")
	}

	return &Currency{Code: code}, nil
}
