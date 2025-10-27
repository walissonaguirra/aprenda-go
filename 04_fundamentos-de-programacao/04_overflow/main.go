/*
Overflow

	Um uint16, por exemplo, vai de 0 a 65535.
	Que acontece se a gente tentar usar 65536?
	Ou se a gente estiver em 65535 e tentar adicionar mais 1?
*/
package main

import (
	"fmt"
)

func main() {
	var x uint16 = 65535
	fmt.Println(x)
	x++
	fmt.Println(x)
	x++
	fmt.Println(x)
}
