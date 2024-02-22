package models

import (
	"github.com/ichbinnichts/events-api/db"
	"github.com/ichbinnichts/events-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password)
		VALUES($1, $2) RETURNING id
	`

	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	hashPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = prepareStatement.QueryRow(u.Email, hashPass).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}
