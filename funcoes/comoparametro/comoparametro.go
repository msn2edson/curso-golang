package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func soma(a, b int) int {
	return a + b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	m := exec(multiplicacao, 3, 4)
	fmt.Println(m)

	s := exec(soma, 1, 2)
	fmt.Println(s)
}
