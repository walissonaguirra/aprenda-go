/*
Funções anônimas

	Anonymous self-executing functions → Funções anônimas auto-executáveis.
	func(p params) { ... }()
	Vamos ver bastante quando falarmos de goroutines.
*/
package main

import "fmt"

func main() {

	func(s string) {
		fmt.Println("Função anonima", s)
	}("!!")

}
