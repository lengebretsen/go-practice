package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type UserModel struct {
	DB *sql.DB
}

type UserRepository interface {
	SelectAllUsers() ([]User, error)
	SelectOneUser(id uuid.UUID) (User, error)
	InsertUser(usr User) (User, error)
	UpdateUser(usr User) (User, error)
	DeleteUser(id uuid.UUID) error
}

func (m UserModel) SelectAllUsers() ([]User, error) {
	var users []User = make([]User, 0)
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}

func (m UserModel) SelectOneUser(id uuid.UUID) (User, error) {
	var user User

	row := m.DB.QueryRow("SELECT * FROM users WHERE Id = UUID_TO_BIN(?)", id)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrModelNotFound
		}
		return user, err
	}
	return user, err
}

func (m UserModel) InsertUser(usr User) (User, error) {
	result, err := m.DB.Exec("INSERT INTO users (id, firstname, lastname) VALUES (UUID_TO_BIN(?), ?, ?)", usr.Id, usr.FirstName, usr.LastName)
	if err != nil {
		return User{}, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if count != 1 {
		return User{}, errors.New(fmt.Sprintf("Invalid number of rows written: %d", count))
	}
	return usr, err
}

func (m UserModel) UpdateUser(usr User) (User, error) {
	result, err := m.DB.Exec("UPDATE users set FirstName = ?, LastName = ? WHERE Id = UUID_TO_BIN(?)", usr.FirstName, usr.LastName, usr.Id)
	if err != nil {
		return User{}, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if count == 0 {
		return User{}, ErrModelNotFound
	}
	return usr, err
}

func (m UserModel) DeleteUser(id uuid.UUID) error {
	//Start new db transaction
	ctx := context.Background()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete address records
	_, err = tx.ExecContext(ctx, "DELETE FROM addresses WHERE UserId = UUID_TO_BIN(?)", id.String())
	if err != nil {
		tx.Rollback()
		return err
	}
	//Delete user record
	res, err := tx.ExecContext(ctx, "DELETE FROM users WHERE Id = UUID_TO_BIN(?)", id.String())
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rows == 0 {
		tx.Rollback()
		return ErrModelNotFound
	}

	//Commit transaction
	err = tx.Commit()
	return err
}
