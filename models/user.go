package models

import (
	"errors"

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

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string

	err := row.Scan(&u.ID, &retrievedPass)

	if err != nil {
		return errors.New("Credentials invalid.")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPass)

	if !passwordIsValid {
		return errors.New("Credentials invalid.")
	}

	return nil
}
