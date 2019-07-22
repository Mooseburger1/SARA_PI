package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//error hanlder
func check(e error) {
	if e != nil {
		fmt.Println("We Panicked!")
		fmt.Println(e)
		log.Fatal(e)
		panic(e)
	}
}

func main() {

	var links []string
	// Initialize configs
	Getconfigs()

	// Get list of image files from DropBox
	p := ListImagesFromDropbox()
	if len(p) < 1 {
		fmt.Println("List of Images From Dropbox is 0\nAre there Images in the /images folder?")
	}

	// Get temporary links from list of image files
	for i := 0; i < len(p); i++ {
		l := GetTemporaryLink(p[i])
		links = append(links, l)
	}
	fmt.Println(links)
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
