package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"G1": 1.1,
			"G2": 1.2,
		},
		"H": {
			"H1": 2.1,
			"H2": 2.2,
		},
		"B": {
			"B1": 3.1,
			"B2": 3.2,
		},
	}

	delete(funcsPorLetra, "H")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
