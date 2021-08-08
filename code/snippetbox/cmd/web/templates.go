package main

import (
	"html/template"
	"path/filepath"
	"time"

	"goFigure/code/snippetbox/pkg/forms"
	"goFigure/code/snippetbox/pkg/models"
)

type templateData struct {
	AuthenticatedUser int
	CSRFToken         string
	CurrentYear       int
	Flash             string
	Form              *forms.Form
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	//init new map for cache
	cache := map[string]*template.Template{}

	//Use .Glob to get slice of all filepaths to target all .page.tmpl's
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	//loop through pages
	for _, page := range pages {
		//extract file name

		name := filepath.Base(page)

		//parse the page template file in to a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//ParseGlob to add layout templates to set "base"
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		//ParseGlob to add 'partial' templates to the set "footer"
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		//add template set to cache using name as key
		cache[name] = ts
	}

	//return map
	return cache, nil
}
