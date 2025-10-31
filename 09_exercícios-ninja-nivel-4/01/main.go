package main

import "fmt"

func main() {
	arr := [5]int{}

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	for _, v := range arr {
		fmt.Println("value:", v)
	}

	fmt.Printf("%T", arr)

}
