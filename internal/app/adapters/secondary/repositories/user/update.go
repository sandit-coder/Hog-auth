package user

import (
	appErr "Hog-auth/internal/app/application/errors"
	"Hog-auth/internal/app/domain/entities"
	"context"
	"fmt"
)

func (r *Repository) Update(ctx context.Context, user *entities.User) error {
	query := "UPDATE users SET user_type  = $1, refresh_token_session_id = %2, email = %3 WHERE id = $4"

	commandTag, err := r.db.Exec(
		ctx,
		query,
		user.UserType(),
		user.RefreshTokenSessionId(),
		user.ID())
	if err != nil {
		return fmt.Errorf("failed to update jwt: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("jwt: %w", appErr.NotFound)
	}

	return nil
}
