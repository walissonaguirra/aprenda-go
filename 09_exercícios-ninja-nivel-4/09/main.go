package main

import "fmt"

func main() {

	g := map[string][]string{}

	g["walisson_aguirra"] = []string{"Estudar", "Programar", "Caminha"}
	g["sara_aguirra"] = []string{"Viajar", "Ler"}
	g["manu_aguirra"] = []string{"Estudar", "Socializar"}

	for key, value := range g {
		fmt.Println("nome:", key, "hobbies:")
		for _, h := range value {
			fmt.Println("-", h)
		}
		fmt.Println("")
	}
}
