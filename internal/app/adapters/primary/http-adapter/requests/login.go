package requests

type Login struct {
	Credential string `json:"credential" validate:"required"`
	Type       string `json:"type" validate:"required"`
}
