/*
Struct

	Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
	Seu nome vem de "structure," ou estrutura.
	Declaração: type x struct { y: z }
	Acesso: x.y
	Exemplo: nome, idade, fumante.
*/
package main

import "fmt"

type person struct {
	name  string
	isDev bool
}

func main() {

	p1 := person{
		name:  "Walisson Aguirra",
		isDev: true,
	}

	p2 := person{
		name:  "Sara Aguirra",
		isDev: false,
	}

	p3 := person{
		"Manu Aguirra",
		true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
