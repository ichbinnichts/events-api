package models

import (
	"time"

	"github.com/ichbinnichts/events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES(?, ?, ?, ?, ?)`

	preperaStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer preperaStatement.Close()

	result, err := preperaStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetAllEvents() []Event {
	return events
}
