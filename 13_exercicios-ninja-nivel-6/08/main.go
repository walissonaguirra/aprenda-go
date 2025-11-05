package main

import "fmt"

func main() {
	hi := fatore()
	hi()
}

func fatore() func() {
	return func() {
		fmt.Println("hi")
	}
}
