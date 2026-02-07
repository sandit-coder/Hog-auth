package user

import (
	appErr "Hog-auth/internal/app/application/errors"
	"Hog-auth/internal/app/domain/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repository) GetById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = %1"

	var user entities.User

	if err := r.db.QueryRow(ctx, query, id).Scan(
		user.ID(),
		user.UserType(),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("jwt: %w", appErr.NotFound)
		}
		return nil, fmt.Errorf("jwt: %w", err)
	}

	return &user, nil
}
