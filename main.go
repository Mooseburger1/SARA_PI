package main

import (
	"fmt"
	"log"
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

// Go routine channels for concurrency
var in = make(chan DropboxHTTPTempLink, 100)
var out = make(chan string, 100)
var status = make(chan int64)

func main() {

	// Initialize configs
	Getconfigs()

	// Get list of image files from DropBox
	go ListImagesFromDropbox(out, status)

	// Get temporary links from list of image files
	go GetTemporaryLink(in, out)

	// Serve up html
	Serve2()
}
