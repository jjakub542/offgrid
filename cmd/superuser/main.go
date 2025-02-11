package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"offgrid/internal/database"
	"offgrid/internal/domain"
	"offgrid/internal/repository"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	var err error
	newUser := domain.User{IsSuperuser: true}

	fmt.Println("Email: ")
	fmt.Scanln(&newUser.Email)
	fmt.Println("Password: ")
	fmt.Scanln(&newUser.Password)

	err = newUser.Validate()
	if err != nil {
		log.Fatal(err)
	}

	hashed_password := sha256.New()
	hashed_password.Write([]byte(newUser.Password))

	newUser.PasswordHash = hex.EncodeToString(hashed_password.Sum(nil))

	repo := repository.New(db)

	err = repo.User.CreateOne(&newUser)

	if err != nil {
		log.Fatal(err)
	}
}
