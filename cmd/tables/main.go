package main

import (
	"offgrid/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	database.InitTables(db, "internal/database/tables.sql")
}
