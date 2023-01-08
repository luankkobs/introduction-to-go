package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

// defer trabsfere a chamada a second para o final da função.
func main() {
	defer second()
	first()
}
