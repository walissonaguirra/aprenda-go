package main

import "fmt"

func main() {
	e := make([]string, 26, 26)

	e = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás",
		"Maranhão", "Mato Grosso", "Mato Grosso do Sul",
		"Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul",
		"Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins"}

	fmt.Println("len:", len(e), "cap:", cap(e))

	for i := 0; i < len(e); i++ {
		fmt.Println((i + 1), e[i])
	}
}
