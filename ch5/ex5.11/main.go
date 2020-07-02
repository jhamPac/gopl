package main

import (
	"fmt"
	"os"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"linear algebra":       {"calculus"},
	"intro to programming": {"data structures"},
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
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Frintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func index(s string, slice []string) (int, error) {
	for i, v := range slice {
		if s == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)
	var visitAll func(items []string, parents []string)

	visitAll = func(items []string, parents []string) {
		for _, k := range items {
			vResolved, seen := resolved[k]
			if seen && !vResolved {
				start, _ := index(k, parents) // ignore error since the key has to be in parents
				err = fmt.Errorf("cycle: %s", strings.Join(append(parents[start:], k), " ->"))
			}
			if !seen {
				resolved[k] = false
				visitAll(m[k], append(parents, k))
				resolved[k] = true
				order = append(order, k)
			}
		}
	}

	for k := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{k}, nil)
	}
	return order, nil
}
