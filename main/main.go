package main

import (
	"fmt"
	"log"

	link "github.com/leoopd/html-link_parser/util"
	build "github.com/leoopd/sitemap-builder/util"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var url string = "https://www.calhoun.io"
	var list link.LinkedList
	found := make(map[string]bool)
	visited := make(map[string]bool)

	html, err := build.GrabHtml(url)
	check(err)

	node, err := link.HtmlParser(html)
	check(err)

	link.TagParser(node, &list)

	links := build.GetLinks(url, list, found)

	for i := 0; i < len(links); i++ {

		if !visited[links[i]] {
			visited[links[i]] = true
			html, err := build.GrabHtml(links[i])
			check(err)

			node, err := link.HtmlParser(html)
			check(err)

			link.TagParser(node, &list)

			links = append(links, build.GetLinks(links[i], list, found)...)
		}
	}
	fmt.Println(links)
}
