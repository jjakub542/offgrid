package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

var (
	database       = os.Getenv("DB_NAME")
	password       = os.Getenv("DB_PASSWORD")
	username       = os.Getenv("DB_USERNAME")
	port           = os.Getenv("DB_PORT")
	host           = os.Getenv("DB_HOST")
	postgresClient *pgx.Conn
)

func New() *pgx.Conn {
	if postgresClient != nil {
		return postgresClient
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
