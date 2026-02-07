package dto

type Verify struct {
	Type       string `json:"type"`
	Code       string `json:"code"`
	Role       string `json:"role"`
	Credential string `json:"credentials"`
}
