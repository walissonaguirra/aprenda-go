/*
Maps: introdução

	Utiliza o formato key:value.
	E.g. nome e telefone
	Performance excelente para lookups.
	map[key]value{ key: value }
	Acesso: m[key]
	Key sem value retorna zero. Isso pode trazer problemas.
	Para verificar: comma ok idiom.

		v, ok := m[key]
		ok é um boolean, true/false

	Na prática: if v, ok := m[key]; ok { }
	Para adicionar um item: m[v] = value
	Maps não tem ordem.
*/
package main

import "fmt"

func main() {
	contacts := map[string]int{
		"walisson": 1234,
		"sara":     5678,
	}

	fmt.Println(contacts)
	fmt.Println(contacts["walisson"])

	// contacts["manu"] = 9999

	// comma ok
	if contact, ok := contacts["manu"]; ok {
		fmt.Println(contact)
	}
}
