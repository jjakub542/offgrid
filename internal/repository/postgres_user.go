package repository

import (
	"offgrid/internal/domain"

	"github.com/jackc/pgx/v5"
)

type postgresUserRepository struct {
	conn *pgx.Conn
}

func NewPostgresUser(c *pgx.Conn) domain.UserRepository {
	return &postgresUserRepository{conn: c}
}

func (p *postgresUserRepository) CreateOne(user *domain.User) error {
	return nil
}

func (p *postgresUserRepository) UpdateOneById(id string, user *domain.User) error {
	return nil
}

func (p *postgresUserRepository) DeleteOneById(id string) error {
	return nil
}

func (p *postgresUserRepository) GetOneByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (p *postgresUserRepository) GetOneById(id string) (domain.User, error) {
	return domain.User{}, nil
}

func (p *postgresUserRepository) GetAll() ([]domain.User, error) {
	return []domain.User{}, nil
}
