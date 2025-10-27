package main

import "fmt"

type numero int

var x numero

func main() {
	fmt.Printf("x: %v -- %T\n", x, x)
	x = 101
	fmt.Println("x:", x)
}
