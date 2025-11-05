package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Àrea do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (q circulo) area() {
	resultado := math.Pi * 2 * q.raio
	fmt.Println("Àrea do circulo:", resultado)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	// x.area()
	// y.area()

	medida(x)
	medida(y)
}
