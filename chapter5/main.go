package main

import "fmt"

func lets() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func main() {
	// Criamos um array de tamanho 5 para armazenar as pontuações e em seguida
	// preenchemos cada elemento com uma nota
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	var total float64 = 0
	// A seguir criamos um laço for para calcular a pontuação total
	for _, value := range x {
		total += value
	}
	// Por fim, dividimos a pontuação total pelo número de elementos para calcular a média.
	fmt.Println(total / float64(len(x)))
}
