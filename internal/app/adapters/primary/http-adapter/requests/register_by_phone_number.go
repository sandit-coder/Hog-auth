package requests

type RegisterByPhoneNumber struct {
	PhoneNumber string `json:"phone_number" validate:"required, e164"`
}
