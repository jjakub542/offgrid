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
	GetById(string) (User, error)
	GetByEmail(string) (User, error)
	Create(*User) error
	Update(*User) error
	DeleteById(string) error
}
