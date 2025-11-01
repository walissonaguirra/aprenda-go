package main

import "fmt"

func main() {

	ss := [][]string{
		[]string{"Walisson", "Aguirra", "Programa"},
		[]string{"Sara", "Aguirra", "Viajar"},
		[]string{"Manu", "Aguirra", "Estudar"},
	}

	fmt.Printf("Nome\t\tSobrenome\tHobby favorito\n")
	for i, v := range ss {
		if i == 0 {
			fmt.Printf("%v\t%v\t\t%v\n", v[0], v[1], v[2])
			continue
		}
		fmt.Printf("%v\t\t%v\t\t%v\n", v[0], v[1], v[2])
	}

}
