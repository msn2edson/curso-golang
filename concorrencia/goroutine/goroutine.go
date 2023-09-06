package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (interacao %d)\n", pessoa, texto, i)
	}
}

func main() {
	//fale("Maria", "Oi", 3)
	//fale("Joao", "Ola", 3)

	//go fale("Maria", "Oi", 30)
	//go fale("Joao", "Ola", 30)

	go fale("Maria", "Oi", 5)
	fale("Joao", "Ola", 2)
}
