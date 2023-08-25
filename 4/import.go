package main

//Importando pacote da propria linguagem go
import "fmt"

type ID int

var (
	a bool    = true
	b int     = 10
	c string  = "Roberto"
	d float64 = 0.5
	e ID      = 1
)

//Printf usa %T para mostrar o tipo ou %v para mostrar o valor
func main() {
	fmt.Printf("O tipo de E é %v", e)
}
