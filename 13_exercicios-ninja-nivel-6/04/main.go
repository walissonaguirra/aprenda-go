package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) show() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func main() {
	pessoa := pessoa{
		"Walisson", "Aguirra", 25,
	}

	pessoa.show()
}
