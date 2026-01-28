package vo

import (
	"fmt"
	"strings"
)

const (
	Seeker   = "seeker"
	Employer = "employer"
)

type Role struct {
	Value string
}

func NewRole(value string) (Role, error) {
	if value == "" {
		return Role{}, fmt.Errorf("value is empty")
	}

	switch strings.ToLower(value) {
	case Seeker:
		return Role{Value: Seeker}, nil
	case Employer:
		return Role{Value: Employer}, nil
	default:
		return Role{}, fmt.Errorf("unknown role")
	}
}
