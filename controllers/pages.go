package controllers

import (
	"html/template"
	"net/http"
)

var (
	indexTemplate = template.Must(
		template.ParseFiles("../views/index.html", "../views/home.html"))
)

type templateParams struct {
	Title       string
	Description string
}

// Home renders the static homepage
func Home(w http.ResponseWriter, r *http.Request) {
	params := templateParams{
		Title:       "Gm Creatives: Application Partners",
		Description: "",
	}
	indexTemplate.Execute(w, params)
}
