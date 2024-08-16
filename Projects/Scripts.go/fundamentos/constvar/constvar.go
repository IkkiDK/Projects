package main

//Várias formas de inicializar uma variável

import (
	"fmt"
	//Esse m na frente posso abreviar o math utilizando apenas o m
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//Colocar := já inicializa declararando ex(var)
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área é:", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false , "oi!"
	fmt.Println(g, h, i)
}