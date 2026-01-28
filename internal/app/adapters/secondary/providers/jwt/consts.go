package jwt

import "time"

const (
	ExpireRefreshToken = time.Hour * 24
	ExpireAccessToken  = time.Minute * 15

	AcccessToken = "acccess_token"
	RefreshToken = "refresh_token"

	SecretKey = "JWT_SECRET_KEY"
)
