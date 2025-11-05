package main

import "fmt"

func main() {
	var hi func()

	hi = func() {
		fmt.Println("hi")
	}

	hi()
}
