package main

import "fmt"

func notaParaConceito(nota float64) (string, float64) {
	switch {
	case nota >= 9 && nota <= 10:
		return "Conceito -> A Nota =", nota
	case nota >= 8 && nota < 9:
		return "Conceito -> B Nota =", nota
	case nota >= 5 && nota < 8:
		return "Conceito -> C Nota =", nota
	case nota >= 3 && nota < 5:
		return "Conceito -> D Nota =", nota
	default:
		return "Conceito -> E Nota =", nota
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.5))
}
