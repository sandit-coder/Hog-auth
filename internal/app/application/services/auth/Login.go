package auth

import (
	"Hog-auth/internal/app/application/dto"
	"context"
	"fmt"
)

func (auth *Auth) Login(loginDto *dto.Login, ctx context.Context) error {

	handler, ok := auth.strategies[loginDto.Type]
	if !ok {
		return fmt.Errorf("unknown auth type")
	}

	err := handler.Initiate(loginDto.Credential, ctx)
	if err != nil {
		return err
	}

	return nil
}
