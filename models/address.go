package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Address struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Street string    `json:"street"`
	City   string    `json:"city"`
	State  string    `json:"state"`
	Zip    string    `json:"zip"`
	Type   string    `json:"type"`
}

type AddressModel struct {
	DB *sql.DB
}

type AddressRepository interface {
	FetchAddresses() ([]Address, error)
	FetchOneAddress(id uuid.UUID) (Address, error)
	InsertAddress(Address) (Address, error)
	UpdateAddress(Address) (Address, error)
	DeleteAddress(id uuid.UUID) error
	FindAddressesByUserId(userId uuid.UUID) ([]Address, error)
}

func (m AddressModel) queryForAddresses(query string, args ...any) ([]Address, error) {
	var addrs []Address = make([]Address, 0)
	rows, err := m.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var addr Address
		if err := rows.Scan(&addr.Id, &addr.UserId, &addr.Street, &addr.City, &addr.State, &addr.Zip, &addr.Type); err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addrs, err
}

func (m AddressModel) FetchAddresses() ([]Address, error) {
	return m.queryForAddresses("SELECT * FROM addresses")
}

func (m AddressModel) FindAddressesByUserId(userId uuid.UUID) ([]Address, error) {
	return m.queryForAddresses("SELECT * FROM addresses WHERE UserId = UUID_TO_BIN(?)", userId)
}

func (m AddressModel) FetchOneAddress(id uuid.UUID) (Address, error) {
	var addr Address

	row := m.DB.QueryRow("SELECT * FROM addresses WHERE Id = UUID_TO_BIN(?)", id)
	err := row.Scan(&addr.Id, &addr.UserId, &addr.Street, &addr.City, &addr.State, &addr.Zip, &addr.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return addr, &ErrModelNotFound{ModelName: "Address", Id: id}
		}
		return addr, err
	}
	return addr, err
}

func (m AddressModel) InsertAddress(addr Address) (Address, error) {
	result, err := m.DB.Exec(
		"INSERT INTO addresses (id, userId, street, city, state, zip, type) VALUES (UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, ?, ?, ?)",
		addr.Id,
		addr.UserId,
		addr.Street,
		addr.City,
		addr.State,
		addr.Zip,
		addr.Type,
	)
	if err != nil {
		return Address{}, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return Address{}, err
	}
	if count != 1 {
		return Address{}, errors.New(fmt.Sprintf("Invalid number of rows written: %d", count))
	}
	return addr, err
}

func (m AddressModel) UpdateAddress(addr Address) (Address, error) {
	result, err := m.DB.Exec(
		"UPDATE addresses set UserId = UUID_TO_BIN(?), Street = ?, City = ?, State = ?, Zip = ?, Type = ? WHERE Id = UUID_TO_BIN(?)",
		addr.UserId,
		addr.State,
		addr.City,
		addr.State,
		addr.Zip,
		addr.Type,
		addr.Id,
	)
	if err != nil {
		return Address{}, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return Address{}, err
	}
	if count == 0 {
		return Address{}, &ErrModelNotFound{ModelName: "Address", Id: addr.Id}
	}
	return addr, err
}

func (m AddressModel) DeleteAddress(id uuid.UUID) error {
	result, err := m.DB.Exec("DELETE FROM addresses WHERE Id = UUID_TO_BIN(?)")
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return &ErrModelNotFound{ModelName: "Address", Id: id}
	}
	return err
}
