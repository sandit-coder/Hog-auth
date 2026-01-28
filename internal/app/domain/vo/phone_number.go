package vo

import (
	"errors"
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

type PhoneNumber struct {
	Value string
}

var e164Regex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

func NewPhoneNumber(value string) (PhoneNumber, error) {
	if value == "" {
		return PhoneNumber{}, errors.New("phone number is empty")
	}

	if !e164Regex.MatchString(value) {
		return PhoneNumber{}, errors.New("phone number must be in E.164 format (e.g. +14155552671)")
	}

	phoneNumber, err := phonenumbers.Parse(value, "")
	if err != nil {
		return PhoneNumber{}, err
	}

	if !phonenumbers.IsValidNumber(phoneNumber) {
		return PhoneNumber{}, errors.New("invalid E.164 phone number")
	}

	normalized := phonenumbers.Format(phoneNumber, phonenumbers.E164)

	return PhoneNumber{normalized}, nil
}
