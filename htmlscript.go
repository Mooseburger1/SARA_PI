package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

//Server starts the server and hanlders for SARA
func Server() {
	port := flag.String("p", "3000", "port to serve on")

	flag.Parse()

	templates = template.Must(template.ParseGlob("./static/*.html"))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.HandleFunc("/home", homePage)
	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
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
