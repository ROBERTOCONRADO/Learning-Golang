package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(5,10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func soma(a, b int) (int, error) {
	if a + b <= 100 {
		return 0, errors.New("No Go as functions são frequentemente usadas para tratar errors, pois não existe try e cat")
	} 
	return a + b, nil
}



