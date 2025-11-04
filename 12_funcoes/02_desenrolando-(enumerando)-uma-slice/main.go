/*
Desenrolando (enumerando) uma slice

	Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
	Exemplos:

	Desenrolando uma slice de ints com como argumento para a função "soma" anterior
	Pode-se passar zero ou mais valores
	O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
*/
package main

import "fmt"

func main() {
	items := []int{1, 2, 3}

	r, l, s := variatica(1, 2, 3, 4, 5, 6)
	fmt.Println(r, l, s)

	r, l, s = variatica(items...)
	fmt.Println(r, l, s)

	r, l, s = variatica()
	fmt.Println(r, l, s)
}

func variatica(n ...int) (int, int, string) {
	fmt.Print("variatica()\t>> ")
	amount := 0
	for _, v := range n {
		amount += v
	}
	return amount, len(n), "Bom dia!"
}
