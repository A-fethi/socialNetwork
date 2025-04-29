package db

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	database *sql.DB
	once     sync.Once
	dbErr    error
)

// GetDB returns a singleton database connection
func GetDB() (*sql.DB, error) {
	once.Do(func() {
		// Open a connection to the SQLite database
		database, dbErr = sql.Open("sqlite3", "./social_network.db")
		if dbErr != nil {
			return
		}

		// Test the connection
		dbErr = database.Ping()
	})

	return database, dbErr
}

// CloseDB closes the database connection
func CloseDB() error {
	if database != nil {
		return database.Close()
	}
	return nil
}