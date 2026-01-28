package validator

import "github.com/go-playground/validator/v10"

type FiberValidator struct {
	validate *validator.Validate
}

func NewFiberValidator(v *validator.Validate) *FiberValidator {
	return &FiberValidator{validate: v}
}

func (fv *FiberValidator) Validate(out any) error {
	return fv.validate.Struct(out)
}
