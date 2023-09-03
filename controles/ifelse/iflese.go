package main

import "fmt"

func imprimirResultados(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultados(7.3)
	imprimirResultados(5.3)
}
