package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(23456))

	// sem sinal - so positivos - uint8 uint16 uint32 uint64
	var b byte = 255 //byte é um número inteiro sem sinal de 8 bits
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal - int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor axio do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	/*
		No caso abaixo podemos utilizar o tipo byte mas
		em geral, é preferível usar runis para representar caracteres
		em Go, pois eles são um tipo mais eficiente e consistente.
		No entanto, bytes podem ser usados em casos em que o tamanho
		de armazenamento é uma preocupação.
	*/
	var i2 rune = 'A' //  um número inteiro de 32 bits que representa um código de ponto Unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais - float32, float64
	var x float32 = 49.99
	fmt.Println("O tipo de x e", reflect.TypeOf(x))
	fmt.Println("O tipo da literal 49.99 e", reflect.TypeOf(49.99)) // default - float64

	// boolean
	bo := true
	fmt.Println("O tipo do bo e", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome e Edson"
	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho da string e", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu
	nome
	e
	Edson`
	fmt.Println("O tamanho da string e", len(s2))

	// char ???
	// var x char = 'B'  -> Erro
	char := 'A'
	fmt.Println("O tipo de char e", reflect.TypeOf(char))
	fmt.Println(char)
}
