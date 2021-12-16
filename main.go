package main

import (
	"fmt"
	"golang-gallery/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTemplate    *views.View
	contactTemplate *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeTemplate.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactTemplate.Render(w, nil))
}

func fourOhFour(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Custom 404")
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Working!!")

	homeTemplate = views.NewView("bootstrap", "views/home.gohtml")
	contactTemplate = views.NewView("bootstrap", "views/contact.gohtml")

	var h http.Handler = http.HandlerFunc(fourOhFour)
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = h

	http.ListenAndServe(":7777", r)
}
