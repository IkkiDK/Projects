package main

import "fmt"

func media(numeros ...float64) float64 {
	resultado := 0.0
	for _, num := range numeros {
		resultado += num
	}
	return resultado / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(6.9, 34.3, 23.5, 6, 1, 22))
}