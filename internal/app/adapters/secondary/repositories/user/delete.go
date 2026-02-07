package user

import (
	appErr "Hog-auth/internal/app/application/errors"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("errors deleting jwt: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("jwt with id %s %w", id, appErr.NotFound)
	}

	return nil
}
