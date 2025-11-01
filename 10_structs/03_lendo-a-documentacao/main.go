/*
Lendo a documentação

É importante se familiarizar com a documentação da linguagem Go.
Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
Veremos:

	ref/spec
	    Já vimos mais da metade dos tipos em Go!
	    Struct types.
	        x, y int
	        anonymous fields
	        promoted fields
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
			name:  "Maria",
			isDev: false,
		},
		age: 12,
	}

	fmt.Println(children)
	fmt.Println(children.person.name)
	fmt.Println(children.name)
}
