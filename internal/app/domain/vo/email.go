package vo

import (
	"fmt"
	"net/mail"
)

type Email struct {
	Value string
}

func NewEmail(value string) (Email, error) {
	address, err := mail.ParseAddress(value)
	if err != nil {
		return Email{}, fmt.Errorf("невалидный email: %w", err)
	}

	return Email{address.Address}, nil
}
