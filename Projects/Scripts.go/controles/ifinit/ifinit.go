package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	//Rand é usada para criar números pseudo-aleatórios
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
func main() {
	if i := numeroAleatorio(); i > 5 { //Tbm suporta essa inicialização no switch
		fmt.Println("Ganhou!!")
	} else {
		fmt.Println("Perdeu!!")
	}
	//Não pode usar o i fora do ifelse (basicamente criado temporariamente)
	//Ex: fmt.Println(i) <- aqui
}
