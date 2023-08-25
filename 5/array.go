package main

import "fmt"

type ID int

var (
	a bool    = true
	b int     = 10
	c string  = "Roberto"
	d float64 = 0.5
	e ID      = 1
)

func main() {
	var meuArray [4]int
	meuArray[0] = 100
	meuArray[1] = 200
	meuArray[2] = 300
	meuArray[3] = 400
	fmt.Println(len(meuArray))             //posição
	fmt.Println(meuArray[len(meuArray)-1]) //valor
//Percorrendo Array
	for i, v := range meuArray {
		fmt.Printf("A posição do indice é %d é o valor é %d\n", i, v)
	}
}
