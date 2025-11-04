/*
Func como expressão

	f := func(p params){ ... }
	f()
*/
package main

import "fmt"

func main() {

	hello := func(s string) string {
		return fmt.Sprintln("Função como expressão", s)
	}

	fmt.Println(hello("!!"))

}
