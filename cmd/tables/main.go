package main

import (
	"context"
	"log"
	"offgrid/internal/database"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.New()
	content, err := os.ReadFile("cmd/tables/tables.sql")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sql := string(content)
	_, err = db.Exec(context.Background(), sql)

	if err != nil {
		log.Fatal(err)
	}
}
