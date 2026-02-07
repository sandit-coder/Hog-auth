package refresh_session

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM refresh_sessions WHERE id = $1"

	commandTag, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting refresh_sessions: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no refresh_sessions found with id %s", id)
	}

	return nil
}
