package main

import "fmt"

func main() {
	// Note-se a ausência de um tamanho entre colchetes. X foi criado com um tamanho igual a 0
	var x []float64

	// Se quiser criar um slice, deve usar a função embutida make
	x = make([]float64, 5)
	fmt.Println(x)
	// Essa instrução acima, cria um slice associada a um array subjacente de float64, de tamanho igual a 5. Os slices
	// são sempre associados a algum array, e, embora não possam, ser maiores que o array, elas podem ser menores.

	x = make([]float64, 5, 10)
	fmt.Println(x)

	arr := [5]float64{1, 2, 3, 4, 5}
	// low é o índice em que o slice começa, e high é o índice em que ele termina (mas não inclui o próprio indice)
	x = arr[0:5]
	fmt.Println(x)
}
