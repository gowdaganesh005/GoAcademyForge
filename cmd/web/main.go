package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gowdaganesh005/GoAcademyForge/internals/models"
)

type application struct {
	infolog  *log.Logger
	errorlog *log.Logger
	test     *models.TestModel
}

func main() {
	addr := flag.String("addr", ":4000", "http network address")
	dsn := flag.String("dsn", "web:pass@/academyforge?parseTime=true", "Mysql dsn")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorlog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infolog:  infolog,
		errorlog: errorlog,
		test:     &models.TestModel{DB: db},
	}

	srv := http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	infolog.Println("Server starting on the port:", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorlog.Println("Error running server:", err)
		return
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
