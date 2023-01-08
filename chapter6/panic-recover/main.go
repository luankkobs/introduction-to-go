package main

import "fmt"

func opa() {
	panic("PANIC")
	str := recover() // isto jamais ocorrerá
	fmt.Println(str)
}

// A chamada panic interrompe imediatamente a execução da função
// Em vez de fazer isso, devemos compor um par com defer.

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
