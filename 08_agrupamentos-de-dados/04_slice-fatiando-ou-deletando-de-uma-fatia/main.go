/*
Slice: fatiando ou deletando de uma fatia

	x[:], x[a:], x[:b], x[a:b]
	"a" é incluso;
	"b" não é.
	Exemplo: cabeça magnética de um disco rígido (relógio, fita).
	    Off-by-one error.
	Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
	É fatiando que se deleta um item de uma slice. Na prática:
	    x := append(x[:i], x[:i]...)
	    Go Playground: https://play.golang.org/p/xK2HwCqvwd
	Exercício: tente acessar todos os itens de uma slice sem utilizar range.
	Solução: https://play.golang.org/p/aUC9qVCobH
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// acessando elementos de um slice
	fmt.Println(slice[3:5])
	fmt.Println(slice[3:])
	fmt.Println(slice[:5])
	fmt.Println(slice[5:len(slice)])

	// Percorrendo uma Slice sem usando range
	for i := 0; i < len(slice); i++ {
		fmt.Println("index:", i, "value:", slice[i])
	}

	// Removendo item de um Slice
	fmt.Println(slice)
	slice = append(slice[:2], slice[9:]...)
	fmt.Println(slice)

}
