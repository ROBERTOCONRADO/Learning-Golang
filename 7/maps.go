package main

import "fmt"

func main() {
	salarios := map[string]int{"roberto": 1000, "Vick": 2000, "Conrado": 3000}
	// fmt.Println(salarios["roberto"])
	delete(salarios, "roberto")
	delete(salarios, "Vick")
	salarios["Rocha"] = 5000
	// fmt.Println(salarios["Rocha"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
		
	}
	//blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}



