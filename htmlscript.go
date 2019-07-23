package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

// ServeUp serves up the html
func ServeUp(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Images []string
	}{Images: []string{"https://morepng.com/public/uploads/preview/4k-hd-cb-backgrounds-11547807360ycumdvi7gz.jpg", "https://morepng.com/public/uploads/preview/4k-hd-cb-backgrounds-11547807360ycumdvi7gz.jpg"}}

	err := templates.ExecuteTemplate(w, "index.html", data)
	check(err)

}

// Serve does blah blah
func Serve() {
	templates = template.Must(template.ParseGlob("site/*.html"))
	router := mux.NewRouter()

	//router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", fs))
	router.HandleFunc("/", ServeUp)

	log.Fatal(http.ListenAndServe(":8000", router))
}
