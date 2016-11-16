//finito in 15 minuti
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	tf     = "~/tedfeed"
	videos = "~/tedfeed/videos"
	thumbs = "~/tedfeed/thumbnails"

	url = "https://www.ted.com/talks/atom"
)

func checkDirectories() {
	if _, err := os.Stat(tf); err != nil {
		fmt.Println("cartella", tf, "non trovata, creazione in corso")
		if err := os.Mkdir(tf, 0755); err != nil {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(videos); err != nil {
		fmt.Println("cartella", videos, "non trovata, creazione in corso")
		if err := os.Mkdir(videos, 0755); err != nil {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(thumbs); err != nil {
		fmt.Println("cartella", thumbs, "non trovata, creazione in corso")
		if err := os.Mkdir(thumbs, 0755); err != nil {
			fmt.Println(err)
		}
	}
}

func httpGetReq() []byte {
	//url fatti in modo hardcoded per semplicita,
	//converebbe comunque costruirli come si deve?
	if resp, err := http.Get(url); err != nil {
		panic("not found")
	} else {

		fmt.Println(resp.Status)
		defer resp.Body.Close()

		var output []byte
		if output, err = ioutil.ReadAll(resp.Body); err != nil {
			panic("not found")
		}

		return output
	}
	return nil
}

func main() {

	checkDirectories()

	output := httpGetReq()

	fmt.Println(string(output))

}
