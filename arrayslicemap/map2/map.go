package main

import "fmt"

func main() {
	funcsESakarios := map[string]float64{
		"Jose joao":      123456.78,
		"Gabriela Silva": 12458.58,
		"Pedro Junior":   45565.55,
	}

	funcsESakarios["Edson"] = 98785.32
	delete(funcsESakarios, "Inexistente")

	for nome, salario := range funcsESakarios {
		fmt.Println(nome, salario)
	}
}
