package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for i, v := range args {
		fmt.Printf("index: %v\tvalue: %v\n", i, v)
	}
}
