/*
Structs embutidos

	Structs dentro de structs dentro de structs.
	Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) e
	tambem um competidor (nome, equipe, pontos).
*/
package main

import "fmt"

type person struct {
	name  string
	isDev bool
}

type children struct {
	person
	age int
}

func main() {

	children := children{
		person: person{
			"Maria",
			false,
		},
		age: 12,
	}

	fmt.Println(children)
}
