package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegnado o endereco da variavel i
	*p++   // desreferenciando (pegando o valor)
	i++

	// go nao tem aritimetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
