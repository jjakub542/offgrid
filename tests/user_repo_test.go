package tests

import (
	"offgrid/internal/domain"
	"offgrid/internal/repository"
	"testing"
)

func TestUserGetOneByEmail(t *testing.T) {
	setupTest(t)
	var err error

	user := &domain.User{
		Email:       "jjakub2d33@gmail.com",
		Password:    "123",
		IsSuperuser: false,
	}

	user.CreatePasswordHash()

	repo := repository.New(TestDB)
	err = repo.User.CreateOne(user)
	if err != nil {
		t.Fatal(err)
	}

	user2, err := repo.User.GetOneByEmail(user.Email)

	if err != nil {
		t.Fatal(err)
	}

	if user2.PasswordHash != user.PasswordHash {
		t.Fail()
	}
}
