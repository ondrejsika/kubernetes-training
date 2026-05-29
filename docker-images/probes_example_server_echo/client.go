package main

import (
	"fmt"
	"io"
	"net/http"
)

func client(path string) error {
	body, err := httpGet("http://127.0.0.1:8000" + path)
	if err != nil {
		return err
	}
	fmt.Print(body)
	return nil
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
