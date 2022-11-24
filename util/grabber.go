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

func GrabHtml(url string) string {
	resp, err := http.Get(url)
	check(err)

	text, err := ioutil.ReadAll(resp.Body)
	return string(text)
}
