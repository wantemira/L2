package main

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func extractLinks(r io.Reader, base *url.URL) []*url.URL {
	var links []*url.URL

	doc, err := html.Parse(r)
	if err != nil {
		return links
	}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if (n.Data == "a" && attr.Key == "href") ||
					(n.Data == "img" && attr.Key == "src") ||
					(n.Data == "script" && attr.Key == "src") ||
					(n.Data == "link" && attr.Key == "href") {

					link, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}

					resolved := base.ResolveReference(link)

					links = append(links, resolved)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}
