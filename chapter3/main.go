package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	x := "Hello, World"
	fmt.Println(x)

	var z string
	z = "Hello World 2"
	fmt.Println(z)

	var y string
	y = "first "
	fmt.Println(y)
	y = y + "second"
	fmt.Println(y)
}
