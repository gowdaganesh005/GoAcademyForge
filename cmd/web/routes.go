package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileserver := http.FileServer(http.Dir("C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\static\\"))

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileserver))
	router.Handler(http.MethodGet, "/about", app.sessionManager.LoadAndSave(http.HandlerFunc(app.about)))
	router.Handler(http.MethodGet, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignup)))
	router.Handler(http.MethodPost, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignupPost)))
	router.Handler(http.MethodGet, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLogin)))
	router.Handler(http.MethodPost, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLoginPost)))
	router.Handler(http.MethodPost, "/user/logout", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLogout)))

	router.Handler(http.MethodGet, "/test/home", app.sessionManager.LoadAndSave(http.HandlerFunc(app.testhome)))
	router.Handler(http.MethodGet, "/test/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.testView)))
	router.Handler(http.MethodGet, "/test/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.testCreate)))
	router.Handler(http.MethodPost, "/test/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.testcreatePost)))

	router.Handler(http.MethodGet, "/reminder/home", app.sessionManager.LoadAndSave(http.HandlerFunc(app.reminderhome)))
	router.Handler(http.MethodGet, "/reminder/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.reminderView)))
	router.Handler(http.MethodGet, "/reminder/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.reminderCreate)))
	router.Handler(http.MethodPost, "/reminder/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.remindercreatePost)))

	router.Handler(http.MethodGet, "/expense/home", app.sessionManager.LoadAndSave(http.HandlerFunc(app.expensehome)))
	router.Handler(http.MethodGet, "/expense/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.expenseView)))
	router.Handler(http.MethodGet, "/expense/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.expenseCreate)))
	router.Handler(http.MethodPost, "/expense/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.expensecreatePost)))

	router.Handler(http.MethodGet, "/attendance/home", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendancehome)))
	router.Handler(http.MethodGet, "/attendance/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendanceView)))
	router.Handler(http.MethodGet, "/attendance/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendanceCreate)))
	router.Handler(http.MethodPost, "/attendance/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendancecreatePost)))
	router.Handler(http.MethodGet, "/attendance/update", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendanceUpdate)))
	router.Handler(http.MethodPost, "/attendance/update", app.sessionManager.LoadAndSave(http.HandlerFunc(app.attendanceUpdatePost)))
	return app.logRequest(router)

}
