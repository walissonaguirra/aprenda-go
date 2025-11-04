/*
Sintaxe

	Qual a utilidade de funções?
	    Abstrair funcionalidade
	    Reutilização de código
	func (receiver) identifier(parameters) (returns) { code }
	A diferença entre parâmetros e argumentos:
	    Funções são definidas com parâmetros
	    Funções são chamadas com argumentos
	Tudo em Go é pass by value.
	    Pass by reference, pass by copy, ... não.
	Parâmetro pode ser ...variádico.
*/
package main

import "fmt"

func main() {
	hello()

	hi("manhã")
	hi("tarde")
	hi("noite")

	result := sum(0.1, 0.2)
	fmt.Println(result)

	r, l, s := variatica(1, 2, 3, 4, 5, 6)
	fmt.Println(r, l, s)
}

func hello() {
	fmt.Print("hello()\t\t>> ")
	fmt.Println("Hello World")
}

func hi(s string) {
	fmt.Print("hi()\t\t>> ")
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

func sum(x, y float32) float32 {
	fmt.Print("sum()\t\t>> ")
	return x + y
}

func variatica(n ...int) (int, int, string) {
	fmt.Print("variatica()\t>> ")
	amount := 0
	for _, v := range n {
		amount += v
	}
	return amount, len(n), "Bom dia!"
}
