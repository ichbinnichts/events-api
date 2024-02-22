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
	query := `INSERT INTO events(name, description, location, datetime, user_id) VALUES($1, $2, $3, $4, $5) RETURNING id`

	prepareStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepareStatement.Close()

	err = prepareStatement.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)
	if err != nil {
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

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
		UPDATE events
		SET name = $1,
		description = $2,
		location = $3,
		datetime = $4

		WHERE id = $5
	`

	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	_, err = prepareStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"

	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	_, err = prepareStatement.Exec(e.ID)

	return err
}
