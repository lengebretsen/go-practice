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
	InsertAddress(Address) (Address, error)
}

func (m AddressModel) FetchAddresses() ([]Address, error) {
	var addrs []Address = make([]Address, 0)
	rows, err := m.DB.Query("SELECT * FROM addresses")
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
