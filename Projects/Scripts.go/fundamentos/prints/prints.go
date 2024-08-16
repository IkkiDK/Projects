package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.1415
	//fmt.Println("O valor de x é " + x)
	//Essa Sprint, ao invéz de imprimir ela retorna o valor
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)
	
	a := 1
	b := 1.9999
	c := false
	d :="Oi!"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	//%V é genérico e geralmente imprime de forma correta
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}