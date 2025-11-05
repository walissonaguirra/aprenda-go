package main

import "fmt"

func main() {
	defer hi()
	hiTwo()
}

func hi() {
	fmt.Println("Hi, Primeiro")
}

func hiTwo() {
	fmt.Println("Oi, Segundo")
}
