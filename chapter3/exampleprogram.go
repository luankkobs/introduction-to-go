package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input) //preenche o input com o número que fornecemos.
	output := input * 2
	fmt.Println(output)
}
