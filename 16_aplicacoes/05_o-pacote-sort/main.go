package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}
