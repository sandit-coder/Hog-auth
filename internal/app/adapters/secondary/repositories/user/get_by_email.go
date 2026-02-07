package user

import (
	appErr "Hog-auth/internal/app/application/errors"
	"Hog-auth/internal/app/domain/entities"
	"Hog-auth/internal/app/domain/vo"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repository) GetByEmail(ctx context.Context, email vo.Email) (*entities.User, error) {
	query := `SELECT u.id, u.user_type
				FROM users u
				JOIN user_credentials c ON c.user_id = u.id
				WHERE u.deleted_at IS NULL
				AND c.deleted_at IS NULL
				AND c.type=$1
				AND c.identifier=$2`

	var user entities.User

	if err := r.db.QueryRow(ctx, query, "email", email.Value).Scan(
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
