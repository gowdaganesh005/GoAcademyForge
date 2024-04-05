package main

import (
	"html/template"
	"path/filepath"

	"github.com/gowdaganesh005/GoAcademyForge/internals/models"
)

type templateData struct {
	Test        *models.Test
	Percentage  float32
	Remarks     string
	Tests       []*models.Test
	CurrentYear int
	Form        any
	Flash       string
	//reminders
	Reminder  *models.Reminder
	Reminders []*models.Reminder

	//expenses
	Expense  *models.Expense
	Expenses []*models.Expense
	Total    float32

	//attendance
	Attendance    *models.Attendance
	Attendances   []*models.Attendance
	AttRemarks    string

	//authentication
	Isauthy bool
}

func newtemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("C:\\Users\\gowda\\Desktop\\GoAcademyForge\\ui\\html\\pages\\*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"C:\\Users\\gowda\\Desktop\\GoAcademyForge\\ui\\partials\\nav.html",
			"C:\\Users\\gowda\\Desktop\\GoAcademyForge\\ui\\partials\\nav2.html",
			"C:\\Users\\gowda\\Desktop\\GoAcademyForge\\ui\\html\\pages\\base.html",
			page,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
