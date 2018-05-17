package main

import (
	"html/template"
	"net/http"

	"github.com/tylergeery/gm-creatives/controllers"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/inquiry/new", controllers.NewInquiry)

	appengine.Main() // Starts the server to receive requests
}

func home(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(
		template.ParseFiles("./views/index.html", "./views/home.html"))

	controllers.Home(w, r, indexTemplate)
}
