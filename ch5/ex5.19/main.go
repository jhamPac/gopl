package main

import "fmt"

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hi"
	}()
	panic("hello from panic")
}

func main() {
	fmt.Println(weird())
}
