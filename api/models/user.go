package models

import (
	"errors"

	"example.com/events-api/db"
	"example.com/events-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = userId
	u.Password = hashPassword
	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, errors.New("could not prepare the query statement")
	}

	row := stmt.QueryRow(email)

	var user User
	err = row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}
		
		users = append(users, user)
	}

	return users, nil
}
