package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileserver := http.FileServer(http.Dir("C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\static\\"))

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileserver))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/user/signup", app.home)
	router.HandlerFunc(http.MethodPost, "/user/signup", app.home)
	router.HandlerFunc(http.MethodGet, "/user/login", app.home)
	router.HandlerFunc(http.MethodPost, "/user/login", app.home)
	router.HandlerFunc(http.MethodPost, "/user/logout", app.home)

	router.HandlerFunc(http.MethodGet, "/test/home", app.testhome)
	router.HandlerFunc(http.MethodGet, "/test/view/:id", app.testView)
	router.HandlerFunc(http.MethodGet, "/test/create", app.testCreate)
	router.HandlerFunc(http.MethodPost, "/test/create", app.testcreatePost)

	router.HandlerFunc(http.MethodGet, "/academy/reminders", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/reminders/id", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/reminders/create", app.home)
	router.HandlerFunc(http.MethodPost, "/academy/reminders/create", app.home)

	router.HandlerFunc(http.MethodGet, "/academy/expenses", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/expenses/id", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/expenses/create", app.home)
	router.HandlerFunc(http.MethodPost, "/academy/expenses/create", app.home)

	router.HandlerFunc(http.MethodGet, "/academy/attendance", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/attendance/id", app.home)
	router.HandlerFunc(http.MethodGet, "/academy/attendance/create", app.home)
	router.HandlerFunc(http.MethodPost, "/academy/attendance/create", app.home)

	return app.logRequest(router)

}
