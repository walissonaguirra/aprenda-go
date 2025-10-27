/*
*
O pacote fmt

	Setup: strings, ints, bools.
	Strings: interpreted string literals vs. raw string literals.
	    Rune literals.
	    Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte.
	Format printing: documentação.
	    Grupo #1: Print -> standard out
	        func Print(a ...interface{}) (n int, err error)
	        func Println(a ...interface{}) (n int, err error)
	        func Printf(format string, a ...interface{}) (n int, err error)
	            Format verbs. (%v %T)
	    Grupo #2: Print -> string, pode ser usado como variável
	        func Sprint(a ...interface{}) string
	        func Sprintf(format string, a ...interface{}) string
	        func Sprintln(a ...interface{}) string
	    Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
	        func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	        func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	        func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
*/
package main

import "fmt"

func main() {

	sterpleted_string_literals := "Olá, tudo bem?\n espero que sim\t Seja bom-vindo. \"Game\""
	raw_string_literals := `"Olá, tudo bem?\n espero que sim\t Seja bom-vindo. \"Game\""`

	fmt.Println("Sterpleted String Literals:", sterpleted_string_literals)
	fmt.Println("Raw String Literals:", raw_string_literals)

	x := "Oi"
	y := "Bom dia!"

	s := fmt.Sprint(x, ", ", y)

	fmt.Print(s)
	fmt.Printf("\ns: %v -- %T", s, s)
}
