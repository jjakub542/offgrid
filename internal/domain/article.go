package domain

import "time"

type Image struct {
	Id         string    `json:"id"`
	Filename   string    `json:"filename"`
	UploadedAt time.Time `json:"uploaded_at"`
	ArticleId  string    `json:"article_id"`
}

type Article struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Public      bool      `json:"public"`
	ImageUrls   []string
}

func (i *Image) GetUrl() string {
	return "/static/uploads/" + i.Filename
}

type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetOneById(string) (*Article, error)
	CreateOne(*Article) error
	UpdateOneById(*Article, string) error
	DeleteOneById(string) error
	AttachImage(*Image, string) error
	GetArticleImages(string) ([]Image, error)
}
