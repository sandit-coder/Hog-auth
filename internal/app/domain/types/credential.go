package types

import (
	"fmt"
	"strings"
)

type CredentialType int

const (
	EmailCredentialType CredentialType = iota
	PhoneCredentialType
)

func CredentialTypeFromString(s string) (CredentialType, error) {
	switch strings.ToLower(s) {
	case "email":
		return EmailCredentialType, nil
	case "phone":
		return PhoneCredentialType, nil
	}
	return 0, fmt.Errorf("invalid registration type: %s", s)
}
