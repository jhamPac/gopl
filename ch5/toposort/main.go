package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			fmt.Println("The item before the if:", item)
			if !seen[item] {
				seen[item] = true
				fmt.Println("the slice being passed in:", m[item])
				visitAll(m[item])
				order = append(order, item)
				fmt.Println("(working backwards now):", item)
				fmt.Println("printing order:", order)
				fmt.Println()
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// sorts the keys so they are in alphabetical order
	sort.Strings(keys)
	fmt.Printf("the keys in order %v:\n", keys)
	visitAll(keys)
	return order
}
