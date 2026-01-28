package user

import (
	"Hog-auth/internal/app/domain/entities"
	domainErr "Hog-auth/internal/app/domain/errors"
	"context"
	"fmt"
)

func (repo *Repository) Update(user *entities.User, ctx context.Context) error {
	query := "UPDATE users SET role = $1, phone_number = %2, email = %3 WHERE id = $4"

	commandTag, err := repo.db.Exec(
		ctx,
		query,
		user.Role(),
		user.PhoneNumber(),
		user.Email(),
		user.ID())
	if err != nil {
		return fmt.Errorf("failed to update auth: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("auth: %w", domainErr.NotFound)
	}

	return nil
}
