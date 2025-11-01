package main

import "fmt"

func main() {
	slice := []float64{1.2, 2.2, 3.4, 0.5, 6.6, 7.2, 8.5, 9.8, 9.9}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])
}
