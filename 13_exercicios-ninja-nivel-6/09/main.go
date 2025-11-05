package main

import "fmt"

func main() {

	handle(hi)

}

func hi() {
	fmt.Println("Ola")
}

func handle(f func()) {
	f()
}
