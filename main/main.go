package main

import (
	"fmt"
	"log"

	build "github.com/leoopd/sitemap-builder/util"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var url string = "https://www.calhoun.io"

	links, _ := build.GetLinks(url)

	for i := 0; i < len(links); i++ {
		fmt.Printf("i: %d, link: %s\n", i, links[i])
	}
	fmt.Println("\n##########\n")
	stringSlice := build.FollowLinks(build.GetLinks(url))

	for i := 0; i < len(links); i++ {
		fmt.Printf("i: %d, link: %s\n", i, stringSlice[i])
	}

}
