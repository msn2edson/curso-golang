package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// utilizacao 1
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	// utilizacao 2
	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"curso golang"}
	fmt.Println(coisa2)
}
