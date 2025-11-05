package main

import "fmt"

func main() {
	i1 := getNumber()
	i2, s := getNumberAndString()

	fmt.Println(i1)
	fmt.Println(i2, s)
}

func getNumber() int {
	return 10
}

func getNumberAndString() (int, string) {
	return 20, "Número vinte"
}
