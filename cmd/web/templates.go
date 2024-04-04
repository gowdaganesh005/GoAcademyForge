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
}

func newtemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\partials\\nav.html",
			"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\base.html",
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