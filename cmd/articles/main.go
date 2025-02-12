package main

import (
	"context"
	"log"
	"offgrid/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	var err error

	db.Exec(context.Background(), `INSERT INTO articles (title, description, content, public)
VALUES 
('Introduction to Go', 'A beginner-friendly guide to Go programming.', 'Go is an open-source programming language designed for simplicity and efficiency.', TRUE);

INSERT INTO articles (title, description, content, public)
VALUES 
('Understanding PostgreSQL Indexes', 'A deep dive into PostgreSQL indexing strategies.', 'Indexes in PostgreSQL help improve query performance by reducing the number of scanned rows.', TRUE);

INSERT INTO articles (title, description, content, public)
VALUES 
('Building REST APIs with Echo', 'A guide to building RESTful APIs using the Echo framework in Go.', 'Echo is a high-performance web framework for Go, designed for rapid development.', TRUE);

INSERT INTO articles (title, description, content, public)
VALUES 
('Concurrency in Go', 'An overview of concurrency patterns in Go.', 'Go provides built-in support for concurrency using goroutines and channels.', FALSE);

INSERT INTO articles (title, description, content, public)
VALUES 
('Database Transactions in Go', 'How to manage database transactions in Go applications.', 'Transactions ensure data integrity by allowing multiple database operations to be executed atomically.', TRUE);
`)

	if err != nil {
		log.Fatal(err)
	}
}
