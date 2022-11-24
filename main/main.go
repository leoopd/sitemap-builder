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

	var list link.LinkedList
	node, err := link.HtmlParser(build.GrabHtml("https://google.com"))
	check(err)

	link.TagParser(node, &list)

	for i := list.Head; i != nil; i = i.Next {
		fmt.Println(i.Link)
	}
}
