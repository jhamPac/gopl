package main

import "fmt"

func comma(s string) string {
	fmt.Println(s)
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	billion := "1000000000"
	b := comma(billion)
	fmt.Printf("Final answer: %s", b)
}
