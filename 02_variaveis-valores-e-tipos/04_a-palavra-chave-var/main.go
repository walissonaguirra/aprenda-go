/*
*
A palavra-chave var

	Variável declarada em um code block é undefined em outro
	Para variáveis com uma abrangência maior, package level scope, utilizamos var
	Funciona em qualquer lugar
	Prestar atenção: chaves, colchetes, parênteses
*/
package main

import "fmt"

var y = 10

func main() {
	z := 20
	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Println(y) // y :: Esta no package-level scope
	fmt.Println(x) // x :: Esta no code-block desta função
}
