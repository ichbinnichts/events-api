package models

import (
	"fmt"
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
	query := `INSERT INTO events(name, description, location, datetime, user_id) VALUES($1, $2, $3, $4, $5) RETURNING id`

	prepareStatement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Prepare error:", err)
		return err
	}
	defer prepareStatement.Close()

	err = prepareStatement.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)
	if err != nil {
		fmt.Println("Exec error:", err)
		return err
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil

}
