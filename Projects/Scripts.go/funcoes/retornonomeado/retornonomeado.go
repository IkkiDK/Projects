package main

import "fmt"

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //Retorno limpo, ele já irá voltar o 'primeiro' e 'segundo'
}

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
}