package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\base.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\home.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\partials\\nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}
func (app *application) testCreate(w http.ResponseWriter, r *http.Request) {
	sub := "dsd"
	ttype := 1
	marks := 10
	total := 20
	id, err := app.test.Insert(sub, ttype, float64(marks), float64(total))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/test/view/%d", id), http.StatusSeeOther)

}
