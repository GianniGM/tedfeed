package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (

	// setting tedfeed path
	tf     = "tedfeed"
	videos = "videos"
	thumbs = "thumbnails"

	// feeds url
	url = "https://www.ted.com/talks/atom"
)

func checkDirectories() {
	home := os.Getenv("HOME")
	videosPath := filepath.Join(home, tf, videos)
	fmt.Println(videosPath)

	// //check tedfeed directory
	// //if not exists creating one
	// if _, err := os.Stat(home + tf); os.IsNotExist(err) {

	// 	fmt.Println("directory", home+tf, "not founded, creating...")
	// 	if err := os.Mkdir(tf, 0755); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	//check tedfeed/videos directory
	//if not exists creating one

	if _, err := os.Stat(videosPath); os.IsNotExist(err) {
		fmt.Println("directory", videos, "not founded, creating...")

		if err := os.MkdirAll(videos, 0755); err != nil {
			fmt.Println(err)
		}

	}

	//check tedfeed/thumbs directory
	//if not exists creating one
	thumbsPath := filepath.Join(home, tf, thumbs)
	fmt.Println(thumbsPath)
	if _, err := os.Stat(thumbsPath); os.IsNotExist(err) {
		fmt.Println("directory", thumbs, "not founded, creating...")

		if err := os.MkdirAll(thumbs, 0755); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {

	checkDirectories()

	//do http GET request to https://www.ted.com/talks/atom
	if resp, err := http.Get(url); err != nil {

		//if connection error
		log.Fatalf("%s", err)

	} else {

		//if succes
		fmt.Println(resp.Status)

		//closing res.Body when finished
		defer resp.Body.Close()

		//read resp contents
		var output []byte
		if output, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Fatalf("%s", err)
		}

		//printing body
		fmt.Println(len(output))
	}
}
