package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type htmlPrettier func(n *html.Node)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=%q", a.Key, a.Val)
			}
			if n.FirstChild != nil {
				fmt.Printf(">\n")
			} else {
				fmt.Printf(" />\n")
			}
			depth++
		} else if n.Type == html.TextNode {
			trimmed := strings.TrimSpace(n.Data)
			if trimmed != "" {
				fmt.Printf("%*s%s\n", depth*2, "", trimmed)
			}
		} else if n.Type == html.CommentNode {
			fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			if n.FirstChild != nil {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode is like a map where you pass in callbacks point free style
func forEachNode(n *html.Node, pre, post htmlPrettier) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
