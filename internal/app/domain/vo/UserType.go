package vo

import (
	"fmt"
	"strings"
)

const (
	Seeker   = "seeker"
	Employer = "employer"
)

type UserType struct {
	Value string
}

func NewUserType(value string) (UserType, error) {
	if value == "" {
		return UserType{}, fmt.Errorf("value is empty")
	}

	switch strings.ToLower(value) {
	case Seeker:
		return UserType{Value: Seeker}, nil
	case Employer:
		return UserType{Value: Employer}, nil
	default:
		return UserType{}, fmt.Errorf("unknown role")
	}
}
