/*
Slice: slice multi-dimensional

	Slices multi-dimensionais são slices que contem slices.
	São como planilhas.
	[][]type
*/
package main

import "fmt"

func main() {
	ss := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}

	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[1][1])
}
