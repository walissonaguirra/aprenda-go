/*
Maps: range & deletando

	Range: for k, v := range map { }
	Reiterando: maps não tem ordem e um range usará uma ordem aleatória.

	delete(map, key)
	Deletar uma key não-existente não retorna erros!
*/
package main

import "fmt"

func main() {
	items := map[int]string{
		1: "Pêra",
		2: "Ova",
		3: "Maçã",
	}

	fmt.Println("key\tvalue")

	for key, value := range items {
		fmt.Println(key, "\t", value)
	}

	delete(items, 2)

	fmt.Println("\nkey\tvalue")

	for key, value := range items {
		fmt.Println(key, "\t", value)
	}
}
