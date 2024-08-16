package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	
	p = &i //Pegando endereço de i
	*p++
	
	//Go não tem aritmética com ponteiros
	//Ex: p++

	fmt.Println(p, *p, i, &i)
}