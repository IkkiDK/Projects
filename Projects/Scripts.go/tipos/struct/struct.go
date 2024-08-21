package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Processador",
		preco:    3.950,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	//Modo simplificado
	produto2 := produto{
		nome:     "Caneta azul",
		preco:    1.50,
		desconto: 0.25,
	}
	fmt.Println(produto2, produto2.precoComDesconto())
}
