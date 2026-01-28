package requests

type PhoneNumberVerification struct {
	Role        string `json:"role" validate:"required"`
	Code        string `json:"code"`
	PhoneNumber string `json:"phone_number" validate:"required, e164"`
}
