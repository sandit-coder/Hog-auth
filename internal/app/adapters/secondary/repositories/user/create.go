package user

import (
	"Hog-auth/internal/app/domain/entities"
	domainErr "Hog-auth/internal/app/domain/errors"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *Repository) Create(user *entities.User, ctx context.Context) (uuid.UUID, error) {
	query := "INSERT INTO users (id, role, phone_number, email) RETURNING id"

	var returnedId uuid.UUID

	if err := repo.db.QueryRow(
		ctx,
		query,
		user.ID(),
		user.Role(),
		user.PhoneNumber(),
		user.Email()).Scan(returnedId); err != nil {

		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == pgerrcode.UniqueViolation {
				return uuid.Nil, fmt.Errorf("auth: %w", domainErr.AlreadyExists)
			}
		}
		return uuid.Nil, fmt.Errorf("auth: %w", err)
	}

	return returnedId, nil
}
