package refresh_session

import (
	"Hog-auth/internal/app/application/ports/repostiries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repostiries.RefreshSession {
	return &Repository{
		db: db,
	}
}
