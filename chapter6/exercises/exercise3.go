package main

import "fmt"

func max(xs ...int) int {
	var max int
	for i, v := range xs {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(max(2, 5, 6, 7, 8, 48, 569))
}
