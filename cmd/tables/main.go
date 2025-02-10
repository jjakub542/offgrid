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
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		is_superuser BOOLEAN
	)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS articles(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		title TEXT UNIQUE,
		description TEXT,
		content TEXT,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		public BOOLEAN,
		author_id UUID NOT NULL,
		FOREIGN KEY (author_id) REFERENCES users(id)
	)`)

	if err != nil {
		log.Fatal(err)
	}
}
