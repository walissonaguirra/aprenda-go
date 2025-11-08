package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	z := pessoa{"walisson", "bandeira", 25}
	fmt.Println(z)
	mudeMe(&z)
	fmt.Println(z)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Sara"
	p.sobrenome = "Malaquias"
}
