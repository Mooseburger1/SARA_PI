package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

// ServeUp serves up the html
func ServeUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
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
	templates = template.Must(template.ParseGlob("static/*.html"))
	router := mux.NewRouter()

	//router.PathPrefix("/site/").Handler(http.StripPrefix("/site/", fs))
	router.HandleFunc("/", ServeUp)

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("listening on port 8000")

}

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
// func (fs FileSystem) Open(path string) (http.File, error) {
// 	f, err := fs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if s.IsDir() {
// 		fmt.Println(path)
// 		index := strings.TrimSuffix(path, "/") + "/index.html"
// 		if _, err := fs.fs.Open(index); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return f, nil
// }

// Serve2 serves the webpage
func Serve2() {
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
