package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)
	
	coisa = 3
	fmt.Println(coisa)
	
	type dinamico interface{}
	
	var coisa2 dinamico = "Oi!"
	fmt.Println(coisa2)
	
	coisa2 = true
	fmt.Println(coisa2)
	
	coisa2 = curso{"Golang: Aprendendo a linguagem do google"}
	fmt.Println(coisa2)
}