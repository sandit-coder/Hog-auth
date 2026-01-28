package dto

type Tokens struct {
	AccessToken  string `json:"access_token" validate:"required, jwt"`
	RefreshToken string `json:"refresh_token" validate:"required,jwt"`
}
