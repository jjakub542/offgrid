package main

import (
	"log"
	"offgrid/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	err := database.InitTables(db, "internal/database/tables.sql")
	if err != nil {
		log.Fatal(err)
	}
}
