package repository

import (
	"offgrid/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	User domain.UserRepository
}

func New(db *pgx.Conn) *Repository {
	return &Repository{User: NewUserRepository(db)}
}
