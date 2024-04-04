package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gowdaganesh005/GoAcademyForge/internals/models"
)

type application struct {
	infolog       *log.Logger
	errorlog      *log.Logger
	test          *models.TestModel
	reminder      *models.RemModel
	attendance    *models.AtModel
	expense       *models.ExpModel
	templateCache map[string]*template.Template
	formdecoder   *form.Decoder
	users         *models.UserModel

	sessionManager *scs.SessionManager
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

	templateCache, err := newtemplateCache()
	if err != nil {
		errorlog.Fatal(err)
	}
	formdecoder := form.NewDecoder()
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	app := &application{
		infolog:        infolog,
		errorlog:       errorlog,
		test:           &models.TestModel{DB: db},
		reminder:       &models.RemModel{DB: db},
		expense:        &models.ExpModel{DB: db},
		attendance:     &models.AtModel{DB: db},
		users:          &models.UserModel{DB: db},
		templateCache:  templateCache,
		formdecoder:    formdecoder,
		sessionManager: sessionManager,
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
