package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gowdaganesh005/GoAcademyForge/internals/models"
	"github.com/gowdaganesh005/GoAcademyForge/internals/validator"
	"github.com/julienschmidt/httprouter"
)

type testcreateForm struct {
	Subject             string  `form:"subject"`
	Testtype            int     `form:"testtype"`
	Marks               float32 `form:"marks"`
	Totalmarks          float32 `form:"totalmarks"`
	validator.Validator `form:"-"`
}
type remindercreateForm struct {
	Title    string `form:"title"`
	Deadline string

	validator.Validator `form:"-"`
}

type expensecreateForm struct {
	Category            string `form:"category"`
	Description         string
	Amount              float32 `form:"amount"`
	Date                string  `form:"date"`
	validator.Validator `form:"-"`
}
type attendanceform struct {
	Subject      string `form:"subject"`
	Attended     int    `form:"attended"`
	Totalclasses int    `form:"totalclasses"`

	validator.Validator `form:"-"`
}
type attendanceupdateForm struct {
	Subject  string `form:"subject"`
	Attended bool
}

type userSignupform struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}
type userLoginform struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
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
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := app.newtemplatedata(r)
	data.Test = test
	data.Percentage = float32((test.Marks / test.Totalmarks) * 100)

	remarks := testRemarks(data.Percentage)
	data.Remarks = remarks
	data.Flash = flash

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
func (app *application) testcreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form testcreateForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.Notless(int(form.Marks), int(form.Totalmarks)), "marks", "Marks Cannot be Greater Than the Total Marks")
	form.CheckField(validator.Notless(0, int(form.Totalmarks)), "totalmarks", "This Field cannot be 0")

	form.CheckField(validator.PermittedInt(form.Testtype, 1, 2, 3, 4), "testtype", "Test Type should be CIE,AAT,SEE")
	if !form.Valid() {
		data := app.newtemplatedata(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "testcreate.html", data)
		return

	}
	id, err := app.test.Insert(form.Subject, form.Testtype, float64(form.Marks), float64(form.Totalmarks))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), "flash", "Test successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/test/view/%d", id), http.StatusSeeOther)

}

func (app *application) testCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = testcreateForm{
		Testtype: 1,
	}

	app.render(w, http.StatusOK, "testcreate.html", data)
}

////////////Reminders

func (app *application) reminderView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notfound(w)

	}
	reminder, err := app.reminder.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrnoRecord) {
			app.notfound(w)

		} else {
			app.serverError(w, err)
			return
		}
	}
	data := app.newtemplatedata(r)
	data.Reminder = reminder

	app.render(w, http.StatusOK, "reminderview.html", data)

}
func (app *application) reminderhome(w http.ResponseWriter, r *http.Request) {
	reminders, err := app.reminder.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newtemplatedata(r)
	data.Reminders = reminders
	app.render(w, http.StatusOK, "reminderhome.html", data)

}
func (app *application) reminderCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = remindercreateForm{
		Deadline: fmt.Sprintf("%+v", time.Now()),
	}

	app.render(w, http.StatusOK, "remindercreate.html", data)
}
func (app *application) remindercreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form remindercreateForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	deadline := r.PostForm.Get("deadline")
	form.Deadline = deadline

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be empty")

	if !form.Valid() {
		data := app.newtemplatedata(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "remindercreate.html", data)
		return

	}
	id, err := app.reminder.Insert(form.Title, form.Deadline)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/reminder/view/%d", id), http.StatusSeeOther)

}

// /////////////////// //expenses
func (app *application) expenseView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notfound(w)

	}
	test, err := app.expense.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrnoRecord) {
			app.notfound(w)

		} else {
			app.serverError(w, err)
			return
		}
	}
	data := app.newtemplatedata(r)
	data.Expense = test

	app.render(w, http.StatusOK, "expenseview.html", data)

}
func (app *application) expensehome(w http.ResponseWriter, r *http.Request) {
	tests, err := app.expense.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newtemplatedata(r)
	data.Expenses = tests
	data.Total, err = app.expense.TotalMonthly()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "expensehome.html", data)

}
func (app *application) expensecreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form expensecreateForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	desc := r.PostForm.Get("description")
	form.Description = desc

	form.CheckField(validator.Notless(0, int(form.Amount)), "amount", "This Field cannot be 0")

	id, err := app.expense.Insert(form.Category, form.Description, form.Amount, form.Date)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/expense/view/%d", id), http.StatusSeeOther)

}
func (app *application) expenseCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = expensecreateForm{
		Category: "Meals",
	}

	app.render(w, http.StatusOK, "expensecreate.html", data)
}

// //////////attendance
func (app *application) attendanceView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notfound(w)

	}
	test, err := app.attendance.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrnoRecord) {
			app.notfound(w)

		} else {
			app.serverError(w, err)
			return
		}
	}
	data := app.newtemplatedata(r)
	data.Attendance = test

	remarks := attendanceRemarks(data.Percentage)
	data.AttRemarks = remarks

	app.render(w, http.StatusOK, "attendanceview.html", data)

}
func (app *application) attendancehome(w http.ResponseWriter, r *http.Request) {
	tests, err := app.attendance.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newtemplatedata(r)
	data.Attendances = tests
	app.render(w, http.StatusOK, "attendancehome.html", data)

}
func (app *application) attendancecreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form attendanceform
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.Notless(int(form.Attended), int(form.Totalclasses)), "attended", "Attended classes Cannot be Greater Than the Total Classes")
	form.CheckField(validator.Notless(0, int(form.Totalclasses)), "totalclasses", "This Field cannot be 0")

	id, err := app.attendance.Insert(form.Subject, form.Attended, form.Totalclasses)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/attendance/view/%d", id), http.StatusSeeOther)

}
func (app *application) attendanceCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = attendanceform{
		Totalclasses: 0,
	}

	app.render(w, http.StatusOK, "attendancecreate.html", data)
}
func (app *application) attendanceUpdatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form attendanceupdateForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	att := r.PostForm.Get("Attended")
	if att == "true" {
		form.Attended = true
	} else {
		form.Attended = false
	}

	id, err := app.attendance.Update(form.Subject, form.Attended)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/attendance/view/%d", id), http.StatusSeeOther)

}
func (app *application) attendanceUpdate(w http.ResponseWriter, r *http.Request) {

	data := app.newtemplatedata(r)
	data.Form = attendanceupdateForm{
		Attended: false,
	}

	app.render(w, http.StatusOK, "attendanceupdate.html", data)
}

/////////user authentication

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = userSignupform{}
	app.render(w, http.StatusOK, "signup.html", data)
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form userSignupform
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")
	if !form.Valid() {
		data := app.newtemplatedata(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
		return
	}
	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")

			data := app.newtemplatedata(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "signup.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	data.Form = userLoginform{}
	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	var form userLoginform
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	if !form.Valid() {
		data := app.newtemplatedata(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		return
	}
	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email or password is incorrect")

			data := app.newtemplatedata(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/test/home", http.StatusSeeOther)
}

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/about", http.StatusSeeOther)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newtemplatedata(r)
	app.render(w, http.StatusOK, "home.html", data)
}
