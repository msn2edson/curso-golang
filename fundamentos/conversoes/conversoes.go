package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 3.2
	y := 2
	fmt.Println(x / float64(y))

	// cuidado
	fmt.Println("Teste " + string(65))

	// int para string - concatenacao
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("65")
	fmt.Println(num - 64)

	// string para bol 1
	b1, _ := strconv.ParseBool("true")
	if b1 {
		fmt.Println("Verdadeiro")
	}

	// string para bol 2
	b2, _ := strconv.ParseBool("1")
	if b2 {
		fmt.Println("Verdadeiro")
	}
}
