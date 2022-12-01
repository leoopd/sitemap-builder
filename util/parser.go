package sitemapBuilder

import (
	"strings"

	link "github.com/leoopd/html-link_parser/util"
)

// Creates a string slice that contains all the links found on the given url and
// a map to keep track of already found links.
func GetLinks(url string) []string {
	var list link.LinkedList
	found := make(map[string]bool)
	var links []string

	node, err := link.HtmlParser(GrabHtml(url))
	check(err)

	link.TagParser(node, &list)

	for i := list.Head; i != nil; i = i.Next {
		if i.Link[0] == 47 {
			if !found[i.Link] {
				found[i.Link] = true
				links = append(links, url+i.Link)
			}
		} else {
			url2 := strings.TrimLeft(url, "htps:/")
			slice := strings.TrimLeft(i.Link, "htps:/")

			if tmp := strings.Split(slice, "/"); tmp[0] == url2 {
				links = append(links, "https://"+slice)
			}
		}
	}
	return links
}

func FollowLinks(links []string) []string {
	visited := make(map[string]bool)

	for i := 0; i < len(links); i++ {
		if !visited[links[i]] {
			visited[links[i]] = true
			links = append(links, GetLinks(links[i])...)
		}
	}
	return links
}
