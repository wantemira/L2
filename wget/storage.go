package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
)

func saveFile(u *url.URL, data []byte, namePath string) {
	filePath := namePath + u.Path

	if filePath == "site" || strings.HasSuffix(filePath, "/") {
		filePath += "index.html"
	}

	dir := path.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir error:", err)
		return
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("write error:", err)
	}
}
