/*
Slice: literal composta

	O que são tipos de dados compostos?

		Wikipedia: Composite_data_type
		Effective Go: Composite literals
		ref/spec: Composite literals

	Uma slice agrupa valores de um único tipo.
	Criando uma slice: literal composta → x := []type{values}
*/
package main

import "fmt"

func main() {
	slice := []int{1, 0, 2}
	fmt.Println(slice)

	slice2 := append(slice, 20)
	fmt.Println(slice2)

	slice[0] = 100
	fmt.Println(slice)
}
