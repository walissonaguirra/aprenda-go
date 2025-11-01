package main

import "fmt"

func main() {
	slice := []float64{1.2, 2.2, 3.4, 0.5, 6.6, 7.2, 8.5, 9.8, 9.9}

	fmt.Println("Type\tValue")

	for _, v := range slice {
		fmt.Printf("%T\t%v\n", v, v)
	}
}
