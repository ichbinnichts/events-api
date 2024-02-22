package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	//Sqllite is a single file to handle sql

	//For use of real databases, instead of 'api.db' put the connection string to your database
	//And and the right driver instead of 'github.com/mattn/go-sqlite3'

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "godev"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic("Could not get database connection.")
	}

	err = DB.Ping()
	if err != nil {
		panic("Could not ping database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table." + err.Error())
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY, 
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime TIMESTAMP WITHOUT TIME ZONE NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id) 
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table." + err.Error())
	}
}
