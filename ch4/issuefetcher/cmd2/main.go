package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jhampac/gopl/ch4/issuefetcher"
)

func main() {
	result, err := issuefetcher.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month old:")

	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 30 {
			formatAndPrint(item)
		}
	}

	fmt.Println("Less than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			formatAndPrint(item)
		}
	}

	fmt.Println("More than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			formatAndPrint(item)
		}
	}
}

func formatAndPrint(i *issuefetcher.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", i.Number, i.User.Login, i.Title)
}
