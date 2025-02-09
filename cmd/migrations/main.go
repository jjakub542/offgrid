package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

var (
	database = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS articles(
		id text,
		title text,
		description text,
		content text,
		created_at text,
		updated_at text,
		public integer,
		CONSTRAINT rid_pkey PRIMARY KEY (id)
	)`)

	if err != nil {
		log.Fatal(err)
	}
}
