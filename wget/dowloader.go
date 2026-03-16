package main

import "io"

func download(link string) ([]byte, string, error) {
	resp, err := client.Get(link)
	if err != nil {
		return nil, "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return body, resp.Header.Get("Content-Type"), nil
}
