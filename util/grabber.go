package sitemapBuilder

import (
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Gets the HTML of the page specified and returns it as a string.
func GrabHtml(url string) (string, error) {
	resp, err := http.Get(url)
	check(err)

	text, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(text), nil
}
