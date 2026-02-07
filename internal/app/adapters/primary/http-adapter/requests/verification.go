package requests

type Verification struct {
	Role       string `json:"role" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Credential string `json:"credential" validate:"required"`
}
