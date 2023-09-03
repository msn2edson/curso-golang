package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos
	turboLibado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // note que 'nome' vem do 'carro'
	f.velocidadeAtual = 0
	f.turboLibado = true

	fmt.Printf("A ferrari %s esta com o turbo ligado? %v\n", f.nome, f.turboLibado)
	fmt.Println(f)
}
