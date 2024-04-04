package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gowdaganesh005/GoAcademyForge/internals/models"
	"github.com/gowdaganesh005/GoAcademyForge/internals/validator"
	"github.com/julienschmidt/httprouter"
)

type testcreateForm struct {
	Subject    string  `form:"subject"`
	Testtype   int     `form:"testtype"`
	Marks      float32 `form:"marks"`
	Totalmarks float32 `form:"totalmarks"`
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	app.render(w, http.StatusOK, "home.html", nil)
}
func (app *application) testcreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form testcreateForm
	err = app.formdecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusUnprocessableEntity)

		return
	}

	form.CheckField(validator.Notless(int(form.Marks), int(form.Totalmarks)), "marks", "Marks Cannot be Greater Than the Total Marks")
	form.CheckField(validator.Notless(int(form.Totalmarks), 0), "totalmarks", "This Field cannot be 0")

	form.CheckField(validator.PermittedInt(form.Testtype, 1, 2, 3, 4), "testtype", "Test Type should be CIE,AAT,SEE")
	if !form.Valid() {
		data := app.newtemplatedata(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return

	}
	id, err := app.test.Insert(form.Subject, form.Testtype, float64(form.Marks), float64(form.Totalmarks))
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/test/view/%d", id), http.StatusSeeOther)

}
func (app *application) testView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notfound(w)

	}
	test, err := app.test.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrnoRecord) {
			app.notfound(w)

		} else {
			app.serverError(w, err)
			return
		}
	}
	data := app.newtemplatedata(r)
	data.Test = test
	data.Percentage = float32((test.Marks / test.Totalmarks) * 100)

	remarks := testRemarks(data.Percentage)
	data.Remarks = remarks

	app.render(w, http.StatusOK, "testview.html", data)

}
func (app *application) testhome(w http.ResponseWriter, r *http.Request) {
	tests, err := app.test.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newtemplatedata(r)
	data.Tests = tests
	app.render(w, http.StatusOK, "testhome.html", data)

}

func (app *application) testCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = testcreateForm{
		Testtype: 1,
	}

	app.render(w, http.StatusOK, "create.html", data)
}
