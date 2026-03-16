package main

import (
	"fmt"
	"net/url"
	"strings"
)

func bfs(u *url.URL, domain string, depth int) {
	if depth == 0 {
		return
	}

	if visited[u.String()] {
		return
	}

	visited[u.String()] = true

	fmt.Println("downloading:", u.String())

	body, contentType, err := download(u.String())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	saveFile(u, body, domain)

	if !strings.Contains(contentType, "text/html") {
		return
	}

	links := extractLinks(strings.NewReader(string(body)), u)

	for _, link := range links {
		if link.Host != domain {
			continue
		}
		bfs(link, domain, depth-1)
	}
}
