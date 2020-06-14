package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	x := strings.Join(os.Args[0:], " ")
	fmt.Println(x)
}
