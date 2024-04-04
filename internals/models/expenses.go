package models

import (
	"database/sql"
	"errors"
)

type Expense struct {
	ID          int
	Category    string
	Description sql.NullString
	Amount      float32
	Date        string
}
type ExpModel struct {
	DB *sql.DB
}

func (t *ExpModel) Insert(category string, description sql.NullString, amount float32, date string) (int, error) {
	stmt := "INSERT INTO expenses (category,description,amount,date) VALUES(?,?,?,TIMESTAMP(?))"
	result, err := t.DB.Exec(stmt, category, description, amount, date)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}
func (t *ExpModel) Get(id int) (*Expense, error) {
	stmt := "SELECT * FROM expenses where id=?"
	row := t.DB.QueryRow(stmt, id)
	te := &Expense{}
	err := row.Scan(&te.ID, &te.Category, &te.Description, &te.Amount, &te.Date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrnoRecord
		} else {
			return nil, err

		}
	}
	return te, nil
}
func (t *ExpModel) Latest() ([]*Expense, error) {
	stmt := "SELECT * FROM expenses ORDER BY id DESC LIMIT 10"

	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tests := []*Expense{}

	for rows.Next() {
		s := &Expense{}
		err := rows.Scan(&s.ID, &s.Category, &s.Description, &s.Amount, &s.Date)
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
