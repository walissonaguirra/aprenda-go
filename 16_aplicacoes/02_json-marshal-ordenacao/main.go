package main

import (
	"encoding/json"
	"fmt"
)

type dev struct {
	Name string
	Lang string
}

func main() {

	dev := dev{
		Name: "Walisson Aguirra",
		Lang: "Go Lang | PHP",
	}

	fmt.Println(dev)
	b, _ := json.Marshal(dev)

	fmt.Println(string(b))

}
