/*
Structs anônimos

	São structs sem identificadores.
	x := struct { name type }{ name: value }
*/
package main

import "fmt"

func main() {

	x := struct {
		name string
		age  int
	}{
		name: "Walisson Aguirra",
		age:  25,
	}

	fmt.Println(x)
}
