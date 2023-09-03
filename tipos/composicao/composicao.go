package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// aqui composicao
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais metodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Ligar turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazer baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
