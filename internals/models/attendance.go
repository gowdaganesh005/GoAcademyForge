package models

import (
	"database/sql"
	"errors"
	"time"
)

type Attendance struct {
	ID           int
	Subject      string
	Attended     int
	TotalClasses int
	Percentage   float32
	UpdatedAt    time.Time
}
type AtModel struct {
	DB *sql.DB
}

func (t *AtModel) Insert(subject string, attended int, totalclasses int) (int, error) {
	stmt := "INSERT INTO attendance (subject,attended,totalclasses,percentage) VALUES(?,?,?,case when totalclasses=0 then 0 else (attended*100)/totalclasses end)"
	result, err := t.DB.Exec(stmt, subject, attended, totalclasses)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}
func (t *AtModel) Get(id int) (*Attendance, error) {
	stmt := "SELECT * FROM attendance where id=?"
	row := t.DB.QueryRow(stmt, id)
	te := &Attendance{}
	err := row.Scan(&te.ID, &te.Subject, &te.Attended, &te.TotalClasses, &te.Percentage, &te.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrnoRecord
		} else {
			return nil, err

		}
	}
	return te, nil
}
func (t *AtModel) Latest() ([]*Attendance, error) {
	stmt := "SELECT * FROM attendance ORDER BY id DESC LIMIT 10"

	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tests := []*Attendance{}

	for rows.Next() {
		s := &Attendance{}
		err := rows.Scan(&s.ID, &s.Subject, &s.Attended, &s.TotalClasses, &s.Percentage, &s.UpdatedAt)
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
func (t *AtModel) Update(subject string, attended bool) (int, error) {
	if attended {
		stmt := "UPDATE attendance SET  attended=attended+1, totalclasses=totalclasses+1,percentage=case when totalclasses=0 then 0 else (attended*100)/totalclasses end, updatedat=CURRENT_TIMESTAMP WHERE subject=?"
		_, err := t.DB.Exec(stmt, subject)
		if err != nil {
			return 0, err
		}

	} else {
		stmt := "UPDATE attendance SET   totalclasses=totalclasses+1,percentage=case when totalclasses=0 then 0 else (attended*100)/totalclasses end, updatedat=CURRENT_TIMESTAMP WHERE subject=?"
		_, err := t.DB.Exec(stmt, subject)
		if err != nil {
			return 0, err
		}
	}
	var id int
	err := t.DB.QueryRow("select id from attendance where subject=?", subject).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
