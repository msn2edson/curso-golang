package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Aprovados")
	for i, nome := range aprovados {
		fmt.Printf("%d) %s\n", i, nome)
	}
}

func main() {
	aprovados := []string{"Aa", "Bb", "Cc", "Ee"}
	imprimirAprovados(aprovados...)
}
