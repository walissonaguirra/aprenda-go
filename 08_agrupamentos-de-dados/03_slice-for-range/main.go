/*
Slice: for range

	Slices:

		Tamanho: len(x)
		Índice específico: x[i] (0-based)

	Para ver todos os itens de uma slice utilizamos o loop for com range.
	Range significa alcance, faixa, extensão.
	For range: for i, v := range x {}
*/
package main

import "fmt"

func main() {
	slice := []string{"Pêra", "Maçã", "Uva", "Batata"}

	for index, value := range slice {
		fmt.Println("index:", index, "value:", value)
	}

	for _, value := range slice {
		fmt.Println("value:", value)
	}
}
