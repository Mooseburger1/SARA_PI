package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

//Server starts the server and hanlders for SARA
func Server() {

	//set default port and make it a command line flag
	port := flag.String("p", "3000", "port to serve on")

	flag.Parse()

	//parse html templates
	templates = template.Must(template.ParseGlob("./static/*.html"))

	//Handler functions
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/images/", ImageHandler)

	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Get list of image files from DropBox
	go ListImagesFromDropbox(out, numImages)

	// Get temporary links from list of image files
	go GetTemporaryLink(in, out)

	go Images()

	err := templates.ExecuteTemplate(w, "load.html", nil)
	check(err)
}

func Images() {

	var links []string

	l := <-numImages

	for link := range in {
		links = append(links, link.Link)
		fmt.Println("I'm compiling links")
		if int64(len(links)) == l {
			break
		}
	}

	data := struct {
		Images []string
	}{Images: links}

	imgLinks <- data
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	data := <-imgLinks
	fmt.Println("I've got the links! Serving template now!")
	err := templates.ExecuteTemplate(w, "index.html", data)
	check(err)
}
