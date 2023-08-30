package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	// abaixo nao funciona
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	fmt.Println(obterResultado(6.2))
}
