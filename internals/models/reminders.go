package models

import (
	"database/sql"
	"errors"
	"time"
)

type Reminder struct {
	ID       int
	Title    string
	Deadline time.Time

	Created time.Time
}
type RemModel struct {
	DB *sql.DB
}

func (t *RemModel) Insert(title string, deadline string) (int, error) {
	stmt := "INSERT INTO reminders (title,deadline) VALUES(?,TIMESTAMP(?))"
	result, err := t.DB.Exec(stmt, title, deadline)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}
func (t *RemModel) Get(id int) (*Reminder, error) {
	stmt := "SELECT * FROM reminders where id=?"
	row := t.DB.QueryRow(stmt, id)
	te := &Reminder{}
	err := row.Scan(&te.ID, &te.Title, &te.Deadline, &te.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrnoRecord
		} else {
			return nil, err

		}
	}
	return te, nil
}
func (t *RemModel) Latest() ([]*Reminder, error) {
	stmt := "SELECT * FROM reminders ORDER BY deadline DESC LIMIT 10"

	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tests := []*Reminder{}

	for rows.Next() {
		s := &Reminder{}
		err := rows.Scan(&s.ID, &s.Title, &s.Deadline, &s.Created)
		if err != nil {
			return nil, err
		}
		tests = append(tests, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tests, nil

}
