/*
*
Valor zero

	Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
	O que é valor zero?
		Quando uma variável é criada em Go — seja com `var`, `new`, `make` ou um literal
		composto — e não recebe um valor inicial, ela é automaticamente inicializada com
		o **valor zero** do seu tipo. Esse valor zero é o padrão definido para cada
		tipo em Go.
	Os zeros:
	    ints: 0
	    floats: 0.0
	    booleans: false
	    strings: ""
	    pointers, functions, interfaces, slices, channels, maps: nil
	Use := sempre que possível.
	Use var para package-level scope.

documentação: https://go.dev/ref/spec#The_zero_value
*/
package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
