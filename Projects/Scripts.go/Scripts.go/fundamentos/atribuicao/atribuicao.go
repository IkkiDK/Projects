package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //Inferência de tipo
	i += 3 //i = i + 3
	i -= 3 //i = i - 3
	i /= 3 //i = i / 2
	i *= 3 //i = i * 2
	i %= 3 //i = i % 2

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	//Pode trocar as variáveis sem precisar de uma aux
	x, y = y, x
	fmt.Println(x, y)
}