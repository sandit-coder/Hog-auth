package auth

import "time"

const (
	UserId  = "user_id"
	Role    = "UserType"
	Revoked = "revoked"
	Active  = "active"

	VerificationCodeLyfetime = time.Minute * 2
)
