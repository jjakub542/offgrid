package tests

import (
	"context"
	"fmt"
	"log"
	"offgrid/internal/database"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TestDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "offgrid_admin", "offgrid123", "localhost", "5432", "offgrid_db_test")
	TestDB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Run tests
	code := m.Run()

	TestDB.Close()
	os.Exit(code)
}

func setupTest(t *testing.T) {
	database.InitTables(TestDB, "../internal/database/tables.sql")
	t.Cleanup(func() {
		database.DropTables(TestDB)
	})
}
