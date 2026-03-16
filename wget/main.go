package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var visited = map[string]bool{}
var client = http.Client{Timeout: 10 * time.Second}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: wget <url> [depth]")
		return
	}

	startURL := os.Args[1]
	depth := 2

	if len(os.Args) > 2 {
		fmt.Sscanf(os.Args[2], "%d", &depth)
	}

	u, err := url.Parse(startURL)
	if err != nil {
		fmt.Println("invalid url:", err)
		return
	}

	bfs(u, u.Host, depth)
}
