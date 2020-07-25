package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"

	mytar "github.com/jhampac/gopl/ch10/arprint/tar"
)

func example() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r, _ := mytar.NewReader(f)
	io.Copy(os.Stdout, r)
}
