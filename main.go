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

func main() {

	Getconfigs()

	p := ListImagesFromDropbox()
	if len(p) < 1 {
		fmt.Println("List of Images From Dropbox is 0\nAre there Images in the /images folder?")
	}

	for i := 0; i < len(p); i++ {
		l := GetTemporaryLink(p[i])
		fmt.Println(l)
	}
}
