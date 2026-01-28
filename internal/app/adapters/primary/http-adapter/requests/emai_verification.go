package requests

type EmailVerification struct {
	Role  string `json:"role" validate:"required"`
	Code  string `json:"code" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
