package main

import "fmt"

type pessoa struct {
	nome                     string
	sobrenome                string
	saboreFavoritosDeSorvete []string
}

func main() {

	p1 := pessoa{
		"Walisson",
		"Aguirra",
		[]string{
			"chocolate",
		},
	}

	p2 := pessoa{
		nome:      "Sara",
		sobrenome: "Aguirra",
		saboreFavoritosDeSorvete: []string{
			"chocolate",
			"baunilha",
		},
	}

	fmt.Println(p1.nome, p1.sobrenome)
	for _, v := range p1.saboreFavoritosDeSorvete {
		fmt.Println("-", v)
	}

	fmt.Println(p2.nome, p2.sobrenome)
	for _, v := range p2.saboreFavoritosDeSorvete {
		fmt.Println("-", v)
	}
}
