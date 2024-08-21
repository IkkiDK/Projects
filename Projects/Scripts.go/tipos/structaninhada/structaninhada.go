package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, quantidade: 2, preco: 12.10}, //Para uma melhor identificação
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("O valor total do pedido é: R$ %.2f", pedido.valorTotal())
}
