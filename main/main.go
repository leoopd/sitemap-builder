package main

import (
	"fmt"
	"log"
	"strings"

	link "github.com/leoopd/html-link_parser/util"
	build "github.com/leoopd/sitemap-builder/util"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	var url string = "https://www.google.com"
	found := make(map[string]bool)
	var links []string
	var list link.LinkedList

	node, err := link.HtmlParser(build.GrabHtml(url))
	check(err)

	link.TagParser(node, &list)

	for i := list.Head; i != nil; i = i.Next {
		if i.Link[0] == 47 {
			if !found[i.Link] {
				found[i.Link] = true
				links = append(links, url+i.Link)
			}
		} else {
			url2 := strings.TrimLeft(url, "https://")
			slice := strings.TrimLeft(i.Link, "https://")

			if tmp := strings.Split(slice, "/"); tmp[0] == url2 {
				links = append(links, "https://"+slice)
			}
		}
	}

	for i := 0; i < len(links); i++ {
		
	}
	fmt.Println(links)
}
