package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//Mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Henrique"
	aprovados[12] = "Mário"
	aprovados[123] = "Josnei"
	fmt.Println(aprovados)
	
	for cpf, nome := range aprovados {
		fmt.Printf(" %s (CPF-> %d)\n", nome, cpf )
	}

	fmt.Println(aprovados[1])
	delete(aprovados, 123)
	fmt.Println(aprovados[123])
}