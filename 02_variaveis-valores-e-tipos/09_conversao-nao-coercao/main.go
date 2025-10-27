/*
*
Conversão, não coerção

	Conversão de tipos é o que soa.
	Em Go não se diz casting, se diz conversion.
	a = int(b)
	ref/spec#Conversions
	Fim da sessão. Parabéns! Dicas, motivação e exercícios.
*/
package main

import "fmt"

type hotdog int

var a hotdog

func main() {

	a = 10
	x := 10

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("x: %v, %T\n", x, x)

	x = int(a)
	a = hotdog(x)

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("a: %v, %T\n", a, a)

}
