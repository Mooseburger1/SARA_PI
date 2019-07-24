package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

// ServeUp serves up the html
func ServeUp(w http.ResponseWriter, r *http.Request) {

	var links []string

	l := <-status

	for link := range in {
		links = append(links, link.Link)
		if int64(len(links)) == l {
			break
		}
	}

	data := struct {
		Images []string
	}{Images: links}
	err := templates.ExecuteTemplate(w, "index.html", data)
	check(err)

}

// Serve does blah blah
func Serve() {
	templates = template.Must(template.ParseGlob("site/*.html"))
	router := mux.NewRouter()

	router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", fs))
	router.HandleFunc("/", ServeUp)

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("listening on port 8000")

}
