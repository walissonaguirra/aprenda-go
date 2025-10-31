/*
Slice: make

	Slices são feitas de arrays.
	Elas são dinâmicas, podem mudar de tamanho.
	Sempre que isso acontece, um novo array é criado e os dados são copiados.
	É conveniente, mas tem um custo computacional.
	Para otimizar as coisas, podemos utilizar make.
	make([]T, len, cap)
	"The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
	len(x), cap(x)
	x[n] onde n > len é out of range. Use append.
	Append > cap modifica o array subjacente.
	pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
	Effective Go.
*/
package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)

	fmt.Println("slice", "\t\t\t\tlength", "\tcapacity")
	fmt.Println(slice, "\t\t\t", len(slice), "\t", cap(slice))

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5

	fmt.Println(slice, "\t\t\t", len(slice), "\t", cap(slice))

	slice = append(slice, 6, 7, 8, 9, 10)

	fmt.Println(slice, "\t\t", len(slice), "\t", cap(slice))

	slice = append(slice, 11)

	fmt.Println(slice, "\t", len(slice), "\t", cap(slice))
}
