package main

import (
	"net/http"

	"github.com/tylergeery/gm-creatives/controllers"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/inquiry/new", controllers.NewInquiry)

	appengine.Main() // Starts the server to receive requests
}
