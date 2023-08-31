package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	roberto := Cliente{
		Nome:  "Roberto",
		Idade: 29,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", roberto.Nome, roberto.Idade, roberto.Ativo)
}
