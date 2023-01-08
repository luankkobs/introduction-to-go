package main

import "fmt"

func zerot(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zerot(&x)
	fmt.Println(x) // x é 0
}
