package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DropboxHTTPListResponse returns a struct of only the "entries"
// key in the returned JSON object from Dropbox. All other key:value
// pairs are ignored from the response body. Should be used
// When Unmarshaling the response body from a "list_folder" request
// See conf.yml for Listlink URL and Dropbox Docs on http responses
// https://dropbox.github.io/dropbox-api-v2-explorer/#files_list_folder
type DropboxHTTPListResponse struct {
	Entries []ImageMetaData `json:"entries"`
}

// ImageMetaData used to further parse the "Entries" array of JSON objects
// returned in the response body of a "list_folder" request.
// See "DropboxHTTPListResponse" for further information
type ImageMetaData struct {
	Path string `json:"path_lower"`
	Name string `json:"name"`
}

// DropboxHTTPTempLink used to unmarshal dropbox http temp link request
type DropboxHTTPTempLink struct {
	Link string `json:"link"`
}

// ListImagesFromDropbox makes a "POST" HTTP Request to Dropbox at the "list_folder"
// API URL. It then returns a slice of image names and paths contained in the Dropbox "/images"
// folder. All images must be contained in an "images" folder within the Dropbox application
// account in order for this function to return the list.
func ListImagesFromDropbox(out chan<- string, status chan<- int64) {
	// Initialize variables
	var result DropboxHTTPListResponse
	client := http.Client{}

	//Construct the request boddy to be passed with the HTTP request
	body, err := json.Marshal(map[string]string{
		"path": "/images"})

	check(err)

	//Setup the HTTP request to be made
	request, err := http.NewRequest("POST", Listlink, bytes.NewBuffer(body))
	request.Header.Set("Authorization", Token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Dropbox-Api-Select-User", "")

	check(err)

	//Make HTTP POST request
	for {
		resp, err := client.Do(request)
		check(err)

		defer resp.Body.Close()

		// Unmarshal the response into result
		json.NewDecoder(resp.Body).Decode(&result)

		status <- int64(len(result.Entries))

		// Iterate through the array of image paths and append them to list
		for i := 0; i < len(result.Entries); i++ {
			out <- result.Entries[i].Path

		}

	}

}

// GetTemporaryLink takes the path of an image in Dropbox and returns a temporary link to that artifact in Dropbox
func GetTemporaryLink(in chan<- DropboxHTTPTempLink, out <-chan string) {
	var path string
	var results DropboxHTTPTempLink

	for path = range out {
		fmt.Println("this is what i'm receiving: ", path)
		body, err := json.Marshal(map[string]string{
			"path": path})

		check(err)

		client := http.Client{}

		request, err := http.NewRequest("POST", Getlink, bytes.NewBuffer(body))
		request.Header.Set("Authorization", Token)
		request.Header.Set("Content-Type", "application/json")

		check(err)

		response, err := client.Do(request)
		check(err)

		json.NewDecoder(response.Body).Decode(&results)

		fmt.Println("this is what i'm sending: ", results.Link)
		in <- results

	}
}
