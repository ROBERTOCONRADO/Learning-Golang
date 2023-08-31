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

func (c Cliente) Desativar() {
	c.Ativo = true
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	roberto := Cliente{
		Nome:  "Roberto",
		Idade: 29,
		Ativo: false,
	}
	roberto.Endereco.Cidade = "Condeúba"

	//executando
	roberto.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Endereco: %s", roberto.Nome, roberto.Idade, roberto.Ativo, roberto.Endereco.Cidade)
}
