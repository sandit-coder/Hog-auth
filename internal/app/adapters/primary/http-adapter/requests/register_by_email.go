package requests

type RegisterByEmail struct {
	Email string `json:"email" validate:"required,email"`
}
