package models

import (
	"database/sql"
	"errors"
	"time"
)

type Test struct {
	ID         int
	Subject    string
	Testtype   int
	Marks      float64
	Totalmarks float64
	Created    time.Time
}
type TestModel struct {
	DB *sql.DB
}

func (t *TestModel) Insert(subject string, ttype int, marks float64, total float64) (int, error) {
	stmt := "INSERT INTO test (subject,testtype,marks,totalmarks) VALUES(?,?,?,?)"
	result, err := t.DB.Exec(stmt, subject, ttype, marks, total)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}
func (t *TestModel) Get(id int) (*Test, error) {
	stmt := "SELECT * FROM test where id=?"
	row := t.DB.QueryRow(stmt, id)
	te := &Test{}
	err := row.Scan(&te.ID, &te.Subject, &te.Testtype, &te.Marks, &te.Totalmarks, &te.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrnoRecord
		} else {
			return nil, err

		}
	}
	return te, nil
}
func (t *TestModel) Latest() ([]*Test, error) {
	stmt := "SELECT * FROM test ORDER BY id DESC LIMIT 10"

	rows,err := t.DB.Query(stmt)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	tests:=[]*Test{}

	for rows.Next(){
		s:=&Test{}
		err:=rows.Scan(&s.ID, &s.Subject, &s.Testtype, &s.Marks, &s.Totalmarks, &s.Created)
		if err != nil {
			return nil, err
		}
		tests=append(tests,s)
	}
	if err=rows.Err();err!=nil{
		return nil,err
	}
	return tests, nil

}
