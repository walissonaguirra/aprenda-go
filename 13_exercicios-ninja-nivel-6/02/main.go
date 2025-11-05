package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	r1 := sumInts(ints...)
	r2 := sumSlice(ints)

	fmt.Println(r1, r2)
}

func sumInts(n ...int) int {
	var a int
	for _, v := range n {
		a += v
	}
	return a
}

func sumSlice(s []int) int {
	var a int
	for _, v := range s {
		a += v
	}
	return a
}
