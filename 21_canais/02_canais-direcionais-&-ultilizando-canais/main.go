package main

import "fmt"

func main() {
	channel := make(chan int)

	go send(channel)
	receive(channel)
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}
