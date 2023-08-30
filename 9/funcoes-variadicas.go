package main

import (
	"fmt"
)

func main() {
	fmt.Println(soma(1, 2, 5, 7, 9, 10, 13, 14, 15, 17, 19, 20, 23, 21, 25))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
