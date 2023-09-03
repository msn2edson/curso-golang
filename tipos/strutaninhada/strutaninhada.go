package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 1.0},
			{2, 2, 2.0},
			{produtoID: 3, qtde: 2, preco: 3.0}, // preferivel
		},
	}
	fmt.Printf("valor total do pedido %.2f", pedido.valorTotal())
}
