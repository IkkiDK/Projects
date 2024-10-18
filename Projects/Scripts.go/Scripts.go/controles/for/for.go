package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { //Não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nPar ou ímpar até 10...")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "-> par ")
		} else {
			fmt.Println(i, "-> ímpar ")
		}
	}

	for {
		//Laço infinito
		fmt.Println("Para sempre... ∞")
		time.Sleep(time.Second)
	}
	//Terá o foreach
}
