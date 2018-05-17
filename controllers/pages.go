package controllers

import (
	"html/template"
	"net/http"
)

type templateParams struct {
	Title       string
	Description string
}

// Home renders the static homepage
func Home(w http.ResponseWriter, r *http.Request, tpl *template.Template) {
	params := templateParams{
		Title:       "Gm Creatives: Application Partners",
		Description: "",
	}
	tpl.Execute(w, params)
}
