package main

//Declaração de escopo global
var (
	b bool
	c int
	d string
	e float32
	f float64
)

func main() {
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}

// Variável de escopo local
func x() {
	var a string
	println(a)
}

//Shorthand
func Shorthand() {
	//var a string = "shorthand"
	a := "Shorthand"
	a = "reatribuindo"
	println(a)
}
