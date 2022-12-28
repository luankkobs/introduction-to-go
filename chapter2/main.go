package main

import "fmt"

func main() {
	// Um espaço também é considerado um caractere, portanto o tamanho da string é 12, e não 11.
	fmt.Println(len("Hello, World"))

	// Strings são indexadas a partir de 0, e não de 1. [1] é o segundo elemento, e não o primeiro.
	fmt.Println("Hello, World"[1])

	// A concatenação usa o mesmo símbolo de adição.
	fmt.Println("Hello, " + "World")
}
