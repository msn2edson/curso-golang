package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// metodo : funcao com receiver (receptor)
func (p produto) precoComDescont() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.10,
	}
	fmt.Println(produto1, produto1.precoComDescont())

	produto2 := produto{"Notebook", 1000.0, 0.1}
	fmt.Println(produto2, produto2.precoComDescont())
}
