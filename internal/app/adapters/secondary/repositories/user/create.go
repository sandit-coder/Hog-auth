package user

import (
	appErr "Hog-auth/internal/app/application/errors"
	"Hog-auth/internal/app/domain/entities"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Create(ctx context.Context, entity *entities.User) (uuid.UUID, error) {
	query := "INSERT INTO users (id, user_type) VALUES ($1,$2) RETURNING user_id"

	var returnedId uuid.UUID

	if err := r.db.QueryRow(
		ctx,
		query,
		entity.ID(), entity.UserType()).Scan(returnedId); err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == pgerrcode.UniqueViolation {
				return uuid.Nil, fmt.Errorf("auth: %w", appErr.AlreadyExists)
			}
		}
		return uuid.Nil, err
	}

	creds := entity.Credentials()
	if len(creds) > 0 {
		copyCount, err := r.db.CopyFrom(
			ctx,
			pgx.Identifier{"user_credentials"},
			[]string{"id", "user_id", "type", "identifier"},
			pgx.CopyFromSlice(len(creds), func(i int) ([]any, error) {
				c := creds[i]
				return []any{c.ID(), entity.ID(), c.CredentialType(), c.Credential()}, nil
			}),
		)
		if err != nil {
			return uuid.Nil, err
		}

		if int(copyCount) != len(creds) {
			return uuid.Nil, fmt.Errorf("expected to insert %d credentials, but inserted %d", len(creds), copyCount)
		}
	}

	return returnedId, nil
}
