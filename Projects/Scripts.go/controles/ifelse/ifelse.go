package main

import "fmt"

func imprimirResultado(nota float64) {
	//Não contém parenteses antes das {}
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}
func main() {
	imprimirResultado(8.5)
	imprimirResultado(3.5)
}