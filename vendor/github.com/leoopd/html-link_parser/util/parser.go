package htmllinkparser

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func check(e error) error {
	if e != nil {
		return e
	}
	return nil
}

func HtmlReader(path string) (string, error) {
	text, err := ioutil.ReadFile(path)
	check(err)
	return string(text), nil
}

func HtmlParser(text string) (*html.Node, error) {

	doc, err := html.Parse(strings.NewReader(text))
	check(err)

	return doc, nil
}

func TagParser(n *html.Node, l *LinkedList) {
	var link string
	var text string
	if n.Type == html.ElementNode && n.Data == "a" {

		if n.Attr[0].Key == "href" {
			link = n.Attr[0].Val
			text = n.FirstChild.Data

			for j := n.FirstChild.NextSibling; j != nil; j = j.NextSibling {
				if j.Type == html.TextNode {
					text += " " + j.Data
				}
			}
			l.Insert(link, text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TagParser(c, l)
	}
}
