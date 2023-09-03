package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicialiados
	aprovados := make(map[int]string)

	aprovados[11111111111] = "Maria"
	aprovados[22222222222] = "Pedro"
	aprovados[33333333333] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[11111111111])
	aprovados[11111111111] = "Edson"
	fmt.Println(aprovados[11111111111])
	delete(aprovados, 11111111111)
	fmt.Println(aprovados[11111111111])
}
