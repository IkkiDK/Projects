package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoas struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//Interfaces são imprementadas implicitamente

func (p pessoas) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoas{"Henrique", "Kasprzak"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 339.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
