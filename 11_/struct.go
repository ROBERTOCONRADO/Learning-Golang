package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	roberto := Cliente{
		Nome:  "Roberto",
		Idade: 29,
		Ativo: true,
	}
	roberto.Endereco.Cidade = "Condeúba"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Endereco: %s", roberto.Nome, roberto.Idade, roberto.Ativo, roberto.Endereco.Cidade)
}
