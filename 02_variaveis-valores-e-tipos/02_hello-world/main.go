/**
Hello world!

    Estrutura básica:
        package main.
        func main: é aqui que tudo começa, é aqui que tudo acaba.
        import.
    Packages:
        Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
        Notação: pacote.Identificador. Exemplo: fmt.Println()
        Documentação: fmt.Println.
    Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
    Variáveis não utilizadas? Não pode: _ nelas.
    ...funções variádicas.
    Lição principal: package main, func main, pacote.Identificador.

documentação: https://pkg.go.dev/fmt#Println
 			  https://go.dev/ref/spec#Blank_identifier
*/

package main

import "fmt"

func main() {

	bytes, errors := fmt.Println("Hello World", "::", "Yahoo IoI")
	fmt.Println("return:", bytes, errors)

	fmt.Println("")

	_, err := fmt.Println("Hello World", "::", "Yahoo IoI")
	fmt.Println("return:", err)

	fmt.Println("")

	x := 16
	y := "strings"
	z := true

	fmt.Println(x, y, z)
}
