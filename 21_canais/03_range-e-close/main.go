package main

import "fmt"

func main() {
	channel := make(chan int)
	go loop(10, channel)

	for v := range channel {
		fmt.Println("channel", v)
	}
}

func loop(l int, c chan<- int) {
	for i := 0; i < l; i++ {
		c <- i
	}
	close(c)
}
