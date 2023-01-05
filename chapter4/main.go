package main

import "fmt"

func main() {
	// Cria uma variável i do tipo int e atribui-lhe o valor 1.
	// i é menor ou igual a 10? Sim: pula para o bloco if.
	for i := 0; i <= 10; i++ {
		// o resto de i / 2 é igual a 0? Não: pula para o bloco else.
		if i%2 == 0 {
			// o resto de i / 2 é igual a 0? Sim: pula para o bloco if.
			// Exibe i, seguido de even, e assim por diante até i ser igual a 11.
			fmt.Println(i, "even")
		} else {
			// exibe i, seguido de odd.
			fmt.Println(i, "odd")
		}
	}
}
