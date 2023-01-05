package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	// O conteúdo do slice1 é copiado para o slice2, mas como slice2 tem espaço somente para dois elementos, apenas os
	// dois elementos do slice1 são copiados.
	fmt.Println(slice1, slice2)
}
