package repository

import (
	"context"
	"offgrid/internal/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresArticleRepository struct {
	db *pgxpool.Pool
}

func NewArticleRepository(c *pgxpool.Pool) domain.ArticleRepository {
	return &postgresArticleRepository{db: c}
}

func (p *postgresArticleRepository) GetAll() ([]domain.Article, error) {
	var articles []domain.Article
	sql := `SELECT * FROM articles;`
	rows, err := p.db.Query(context.Background(), sql)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article domain.Article
		if err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Public,
		); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}

func (p *postgresArticleRepository) GetOneById(id string) (*domain.Article, error) {
	var article domain.Article
	sql := `SELECT * FROM articles WHERE id=$1;`
	err := p.db.QueryRow(context.Background(), sql, id).Scan(
		&article.Id,
		&article.Title,
		&article.Description,
		&article.Content,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.Public,
	)
	return &article, err
}

func (p *postgresArticleRepository) CreateOne(a *domain.Article) error {
	sql := `INSERT INTO articles (title, description, content, public)
	VALUES ($1, $2, $3, $4)`
	_, err := p.db.Exec(context.Background(), sql, a.Title, a.Description, a.Content, a.Public)
	return err
}

func (p *postgresArticleRepository) UpdateOneById(a *domain.Article, id string) error {
	sql := `UPDATE articles SET
	title=$1, description=$2, content=$3, updated_at=$4, public=$5
	WHERE id=$6`
	_, err := p.db.Exec(context.Background(), sql, a.Title, a.Description, a.Content, time.Now(), a.Public, id)
	return err
}

func (p *postgresArticleRepository) DeleteOneById(id string) error {
	sql := `DELETE FROM articles WHERE id=$1`
	_, err := p.db.Exec(context.Background(), sql, id)
	return err
}
