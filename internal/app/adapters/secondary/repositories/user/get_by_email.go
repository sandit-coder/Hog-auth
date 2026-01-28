package user

import (
	"Hog-auth/internal/app/domain/entities"
	domainErr "Hog-auth/internal/app/domain/errors"
	"Hog-auth/internal/app/domain/vo"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (repo *Repository) GetByEmail(email vo.Email, ctx context.Context) (*entities.User, error) {
	query := "SELECT * FROM users WHERE email = %1"

	var user entities.User

	if err := repo.db.QueryRow(ctx, query, email.Value).Scan(
		user.ID(),
		user.Role(),
		user.Email(),
		user.PhoneNumber(),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("auth: %w", domainErr.NotFound)
		}
		return nil, fmt.Errorf("auth: %w", err)
	}

	return &user, nil
}
