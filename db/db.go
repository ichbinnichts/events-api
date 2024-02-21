package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	//Sqllite is a single file to handle sql

	//For use of real databases, instead of 'api.db' put the connection string to your database
	//And and the right driver instead of 'github.com/mattn/go-sqlite3'

	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not get database connection.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
