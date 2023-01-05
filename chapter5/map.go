package main

import "fmt"

func main() {
	// Map de strings
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	// Map de int
	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1], len(y))
	delete(y, 1)
}
