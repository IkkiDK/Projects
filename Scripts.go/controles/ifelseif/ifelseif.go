package main

import "fmt"

func notaParaConceito(nota float64) (string, float64) {
	if nota >= 9 && nota <= 10 {
	return "Conceito -> A Nota =", nota
	} else if nota >= 8 && nota < 9 {
		return "Conceito -> B Nota =", nota
	} else if nota >= 5 && nota < 8 {
		return "Conceito -> C Nota =", nota
	} else if nota >= 3 && nota < 5 {
		return "Conceito -> D Nota =", nota
	} else {
		return "Conceito -> E Nota =", nota
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.5))
}