package user

import (
	"Hog-auth/internal/app/domain/entities"
	domainErr "Hog-auth/internal/app/domain/errors"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (repo *Repository) GetById(id uuid.UUID, ctx context.Context) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = %1"

	var user entities.User

	if err := repo.db.QueryRow(ctx, query, id).Scan(
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
