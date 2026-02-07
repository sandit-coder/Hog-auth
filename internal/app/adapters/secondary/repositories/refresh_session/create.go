package refresh_session

import (
	"Hog-auth/internal/app/domain/entities"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *Repository) Create(ctx context.Context, entity *entities.RefreshSession) error {
	query := `INSERT INTO refresh_sessions 
    			(id, user_ID, refresh_token_hash, user_type, revoked)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`

	_, err := repo.db.Exec(
		ctx,
		query,
		entity.ID(),
		entity.UserId(),
		entity.RefreshTokenHash(),
		entity.UserType(),
		entity.Revoked(),
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("refresh_sessions found with id %s", entity.ID())
			}
		}

		return err
	}

	return nil
}
