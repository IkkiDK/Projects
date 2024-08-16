package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //Compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i, numero)
	}

	for _, numero := range numeros { //Pode usar sem o índice
		fmt.Println(numero)
	}
}
