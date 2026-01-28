package user

import (
	domainErr "Hog-auth/internal/app/domain/errors"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *Repository) Delete(id uuid.UUID, ctx context.Context) error {
	query := "DELETE FROM users WHERE id = $1"

	commandTag, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("errors deleting auth: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("auth with id %s %w", id, domainErr.NotFound)
	}

	return nil
}
