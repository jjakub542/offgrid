package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"offgrid/internal/database"
	"offgrid/internal/domain"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	var err error
	newUser := domain.User{}

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

	sql := `INSERT INTO users (email, password_hash, is_superuser) VALUES ($1, $2, $3) RETURNING id, created_at`
	row := db.QueryRow(context.Background(), sql, newUser.Email, newUser.PasswordHash, true)

	user := domain.User{}
	err = row.Scan(&user.Id, &user.CreatedAt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Id)

}
