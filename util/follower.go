package sitemapBuilder

import (
	"fmt"

	link "github.com/leoopd/html-link_parser/util"
)

func FollowLinks(links []string, found map[string]bool, visited map[string]bool, list link.LinkedList) []string {
	for i := 0; i < 10; i++ {
		if !visited[links[i]] {
			visited[links[i]] = true
			newLinks, found := GetLinks(links[i])
			for j := 0; j < len(newLinks); j++ {

				fmt.Println(newLinks[j])
				if !found[newLinks[j]] {
					fmt.Println("new")
					found[newLinks[j]] = true
					links = append(links, newLinks[j])
				}
			}
		}
	}
	return links
}

func checkingFollower (url string) bool {

}

func addingFollower (url string) {

}