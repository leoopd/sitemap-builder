package sitemapBuilder

import (
	"strings"

	link "github.com/leoopd/html-link_parser/util"
)

func cleanLink(url string) string {
	url = strings.TrimLeft(url, "https://")
	return url
}

// Creates a string slice that contains all the links found on the given url and
// a map to keep track of already found links.
func GetLinks(url string, list link.LinkedList, found map[string]bool) []string {
	var links []string

	for i := list.Head; i != nil; i = i.Next {
		if i.Link[0] == 47 {
			if !found[i.Link] {
				found[i.Link] = true
				links = append(links, url+i.Link)
			}
		} else {
			tmpUrl := cleanLink(url)
			tmpSlice := cleanLink(i.Link)

			if tmp := strings.Split(tmpSlice, "/"); tmp[0] == tmpUrl {
				found[tmpSlice] = true
				links = append(links, "https://"+tmpSlice)
			}
		}
	}
	return links
}

// func FollowLinks(links []string, found map[string]bool, visited map[string]bool, list link.LinkedList) []string {
// 	for i := 0; i < 10; i++ {
// 		if !visited[links[i]] {
// 			visited[links[i]] = true
// 			newLinks, found := GetLinks(links[i])
// 			for j := 0; j < len(newLinks); j++ {

// 				fmt.Println(newLinks[j])
// 				if !found[newLinks[j]] {
// 					fmt.Println("new")
// 					found[newLinks[j]] = true
// 					links = append(links, newLinks[j])
// 				}
// 			}
// 		}
// 	}
// 	return links
// }
