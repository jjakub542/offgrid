package repository

import (
	"context"
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
	sql := `INSERT INTO users (email, password_hash, is_superuser) VALUES ($1, $2, $3)`
	_, err := p.conn.Exec(context.Background(), sql, user.Email, user.PasswordHash, user.IsSuperuser)
	return err
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
