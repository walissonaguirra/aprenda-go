/*
Slice: anexando a uma slice

	Effective Go: append (package builtin)
	x = append(slice, ...values)
	x = append(slice, slice...)
	Todd: unfurl → desdobrar, desenrolar
	Nome oficial: enumeration
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	sliceTwo := []int{5, 6, 7, 8, 9, 10}

	fmt.Println(slice)
	slice = append(slice, sliceTwo...)
	fmt.Println(slice)
}
