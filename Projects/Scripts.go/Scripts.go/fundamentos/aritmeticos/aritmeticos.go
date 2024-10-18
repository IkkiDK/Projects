package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplcação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b)

	//Bitwise
	fmt.Println("E =>", a&b)   //11 & 12 = 10
	fmt.Println("Ou =>", a|b)  //11 | 10 = 11
	fmt.Println("Xor =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//Outras opções usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(float64(c), float64(d)))
	fmt.Println("Exponenciação =>", math.Pow(float64(c), float64(d)))
}
