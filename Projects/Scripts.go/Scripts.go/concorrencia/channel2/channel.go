package main

import (
	"fmt"
	"time"
)

//Channel (canal) - é a forma de comunicação entre goroutines
//Channel é um tipo ex: int, float...

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //Enviando dados para o channel

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c //Recebendo os dados do channel
	fmt.Println(a, b)

	fmt.Println(<-c)
}
