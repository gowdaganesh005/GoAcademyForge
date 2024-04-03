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
	stmt:="SELECT * FROM test where id=?"
	row:=t.DB.QueryRow(stmt,id)
	te:=&Test{}
	err:=row.Scan(&te.ID,&te.Subject,&te.Testtype,&te.Marks,&te.Totalmarks)
if err != nil {
	if errors.Is(err,sql.ErrNoRows) {
		return nil,ErrnoRecord
	}else{
		return nil,err

	}
}
return te,nil
}
func (t *TestModel) Latest() ([]*Test, error) {
	return nil, nil

}
