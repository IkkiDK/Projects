package main

import (
	"fmt"

	"github.com/IkkiDK/html"
)

func encamihar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - misturar (mensagens) em 1 canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encamihar(entrada1, c)
	go encamihar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.google.com.br", "https://www.youtube.com.br"),
		html.Titulo("https://www.amazon.com.br", "https://www.instagram.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
