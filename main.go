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
var in = make(chan DropboxHTTPTempLink)
var out = make(chan string)
var numImages = make(chan int64)
var imgLinks = make(chan struct {
	Images []string
})

func main() {

	// Initialize configs
	Getconfigs()

	// Start the Server
	Server()
}
