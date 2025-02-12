package repository

import (
	"offgrid/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User    domain.UserRepository
	Article domain.ArticleRepository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Article: NewArticleRepository(db),
	}
}
