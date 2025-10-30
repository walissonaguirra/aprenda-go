/*
Iota

	golang.org/ref/spec
	Numa declaração de constantes, o identificador iota representa números sequenciais.
	Na prática.

		iota, iota + 1, a = iota b c, reinicia em cada const, _
*/
package main

import (
	"fmt"
)

const (
	a = iota
	_
	c
	d
)

const (
	e = iota * 12
	_
	g
	i
)

func main() {
	fmt.Println(a, c, d)
	fmt.Println(e, g, i)
}
