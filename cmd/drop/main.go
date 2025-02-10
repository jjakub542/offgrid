package main

import (
	"context"
	"log"
	"offgrid/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	var err error

	_, err = db.Exec(context.Background(), `
	DROP TABLE articles;
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
	DROP TABLE users;
	`)

	if err != nil {
		log.Fatal(err)
	}
}
