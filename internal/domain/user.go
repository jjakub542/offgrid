package domain

type User struct {
	Id             string
	Email          string
	Password       string
	IsSuperuser    bool
	DateRegistered string
	Active         bool
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetOneById(string) (User, error)
	GetOneByEmail(string) (User, error)
	CreateOne(*User) error
	UpdateOneById(string, *User) error
	DeleteOneById(string) error
}
