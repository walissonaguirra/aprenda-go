/*
Métodos

	Um método é uma função anexada a um tipo.
	Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
	Pode-se anexar uma função a um tipo utilizando seu receiver.
	Utilização: valor.método()
	Exemplo: o tipo "pessoa" pode ter um método oibomdia()
*/
package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) hi(prefixo string) {
	fmt.Println(prefixo, p.nome, p.sobrenome, "diz Bom dia!!")
}

func main() {

	p := pessoa{
		nome:      "Walisson",
		sobrenome: "Aguirra",
	}

	p.hi("Oi,")

}
