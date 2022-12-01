package sitemapBuilder

import (
	"fmt"
	"strings"

	link "github.com/leoopd/html-link_parser/util"
)

// Creates a string slice that contains all the links found on the given url and
// a map to keep track of already found links.
func GetLinks(url string) ([]string, map[string]bool) {
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
				found[slice] = true
				links = append(links, "https://"+slice)
			}
		}
	}
	return links, found
}

func FollowLinks(links []string, found map[string]bool) []string {
	fmt.Println(found)
	visited := make(map[string]bool)

	for i := 0; i < len(links); i++ {
		if !visited[links[i]] {
			visited[links[i]] = true
			newLinks, found := GetLinks(links[i])
			for j := 0; j < len(newLinks); j++ {

				if !found[newLinks[j]] {
					found[newLinks[j]] = true
					links = append(links, newLinks[j])
				}
			}
		}
	}
	return links
}
