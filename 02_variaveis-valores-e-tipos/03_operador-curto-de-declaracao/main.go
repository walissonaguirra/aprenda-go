/**
Operador curto de declaração

    := parece uma marmota (gopher) ou o punisher.
    Uso:
        Tipagem automática
        Só pode repetir se houverem variáveis novas
        != do assignment operator (operador de atribuição)
        Só funciona dentro de codeblocks
    Terminologia:
        keywords (palavras-chave) são termos reservados
        operadores, operandos
        statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
        expressão -> qualquer coisa que "produz um resultado"
        scope (abrangência)
            package-level scope
    Lição principal:
        := utilizado pra criar novas variáveis, dentro de code blocks
        = para atribuir valores a variáveis já existentes

documentação: https://go.dev/ref/spec#Keywords
*/

package main

import "fmt"

func main() {
	x := 10
	y := "Hello"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30.03
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// 10 < 8 :: Isso é um Expressão (expression)
	// Uma expressão é qualquer coisa que gera um resultado
	// mas não produz uma ação. Ex.

	// 10 < 8 :: Isso sozinho no codigo gerar um resuldado
	// mas não produz/gera nenhum ação.

	// obs: 10 < 8 :: Isso vai gerar um evalued but not used
	// Avaliação de expressões aritméticas
	// Avaliação de funções
	// Avaliação de condições

	a := 10 < 8

	// a := , fmt.Println(a) :: Isso são stantement, pois isso gera
	// uma ação no programa.

	// a := :: Isso gerar uma ação de declarar e atribuir um valor a
	// uma variavel

	// fnt.Println() :: Isso gerar à ação de executar o função

	fmt.Println("10 < 8:", a)
}
