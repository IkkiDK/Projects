package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtda int) {
	for i := 0; i < qtda; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n]", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Josnei", "Tudo bom?", 2)
	//fale("Mario", "Positivo.", 3)

	go fale("Josnei", "Tudo bom?", 3)
	go fale("Mario", "Positivo.", 3)
	time.Sleep(time.Second * 10)
	fmt.Println("Fim!")
}
