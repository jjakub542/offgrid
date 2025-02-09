package domain

type Image struct {
	Id          string
	Name        string
	Description string
	Article     Article
}

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedAt   string
	UpdatedAt   string
	Visible     bool
}

func (i *Image) GetPath() string {
	return "/media/" + i.Article.Id + "/" + i.Id
}

type ArticleRepository interface {
	GetById(string) (Article, error)
	GetByEmail(string) (Article, error)
	Create(*Article) error
	Update(*Article) error
	DeleteById(string) error
}
