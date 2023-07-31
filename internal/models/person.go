package models

import (
	"context"
	"time"
)

type Person struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (m *DBModel) GetPersons() ([]*Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var persons []*Person

	query := `SELECT person_id, firstname, lastname, description, created_at, updated_at FROM persons`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Person
		err = rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &p)
	}
	return persons, nil
}
