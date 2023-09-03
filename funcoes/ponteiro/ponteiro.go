package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao: um ponteiro eh representado por um +
func inc2(n *int) {
	// revisao: * eh usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisao: & usado para obter o endereco da variavel
	inc2(&n)
	fmt.Println(n) // por referencia

}
