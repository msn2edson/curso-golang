package main

import (
	"fmt"
)

// PersonBuilder é uma estrutura que permite construir uma pessoa passo a passo.
type PersonBuilder struct {
	person Person
}

// Person é a estrutura que representa uma pessoa.
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

// Novo constrói uma nova instância de PersonBuilder.
func Novo() *PersonBuilder {
	return &PersonBuilder{}
}

// ComFirstName define o primeiro nome da pessoa.
func (b *PersonBuilder) ComFirstName(firstName string) *PersonBuilder {
	b.person.FirstName = firstName
	return b
}

// ComLastName define o último nome da pessoa.
func (b *PersonBuilder) ComLastName(lastName string) *PersonBuilder {
	b.person.LastName = lastName
	return b
}

// ComAge define a idade da pessoa.
func (b *PersonBuilder) ComAge(age int) *PersonBuilder {
	b.person.Age = age
	return b
}

// ComAddress define o endereço da pessoa.
func (b *PersonBuilder) ComAddress(address string) *PersonBuilder {
	b.person.Address = address
	return b
}

// Construir cria a pessoa com os atributos definidos.
func (b *PersonBuilder) Construir() Person {
	return b.person
}

func main() {
	// Usando o builder para criar uma pessoa.
	pessoa := Novo().
		ComFirstName("João").
		ComLastName("Silva").
		ComAge(30).
		ComAddress("123 Rua Principal").
		Construir()

	// Imprimindo a pessoa criada.
	fmt.Printf("Nome: %s %s\nIdade: %d\nEndereço: %s\n", pessoa.FirstName, pessoa.LastName, pessoa.Age, pessoa.Address)
}
