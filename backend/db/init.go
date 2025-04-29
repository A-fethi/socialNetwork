package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initdb initializes the database and runs the schema
func Initdb() {
	// Open the SQLite database (will create if not exists)
	var err error
	DB, err = sql.Open("sqlite3", "./db/db.db")
	if err != nil {
		panic(err)
	}

	// Ensure the database is available
	if err := DB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	sch, err := os.ReadFile("./db/schema.sql")

	if err != nil {
		panic(fmt.Sprintf("Failed to read schema.sql: %v", err))
	}
	_, err = DB.Exec(string(sch))

	if err != nil {
		panic(fmt.Sprintf("Failed to execute schema: %v", err))
	}

	fmt.Println("Database initialized successfully")
}