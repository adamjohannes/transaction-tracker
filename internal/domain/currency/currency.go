package currency

import (
	"errors"
	"regexp"
)

type Currency struct {
	code string `json:"code"`
}

func New(code string) (*Currency, error) {
	if len(code) != 3 {
		return nil, errors.New("currency must be 3 characters long")
	}

	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !re.MatchString(code) {
		return nil, errors.New("currency code must only contain alphabetic characters")
	}

	return &Currency{code: code}, nil
}
