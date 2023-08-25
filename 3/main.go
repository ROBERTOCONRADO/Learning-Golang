package main 

//Possibilita fazer a criação da minha propria tipagem
type ID int

var (
	a ID = 1
)

func main() {
	println(a)
}