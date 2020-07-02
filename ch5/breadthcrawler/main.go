package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/jhampac/gopl/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var originalHost string

func save(rawURL string) error {
	url, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}
	if originalHost == "" {
		originalHost = url.Host
	}
	if originalHost != url.Host {
		return nil
	}

	dir := url.Host
	var filename string
	if filepath.Ext(filename) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func crawl(url string) []string {
	fmt.Println(url)
	err := save(url)
	if err != nil {
		log.Printf(`can't cache %s: %s`, url, err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Printf(`can't extract links from %s : %s`, url, err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
