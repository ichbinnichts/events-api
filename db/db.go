package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not get database connection.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
