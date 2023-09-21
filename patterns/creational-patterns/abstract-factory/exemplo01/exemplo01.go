package main

import(
  "fmt"
)

// AbstractFactory é uma interface para criar famílias de objetos relacionados.
type AbstractFactory interface {
  CreateShoe() Shoe
  CreateShirt() Shirt
}

// Shoe é uma interface para objetos de sapato.
type Shoe interface {
  SayName() string
}

// AdidasShoe é uma implementação concreta de Shoe.
type AdidasShoe struct{}

func (a AdidasShoe) SayName() string {
  return "Adidas Shoe"
}

// NikeShoe é uma implementação concreta de Shoe.
type NikeShoe struct{}

func (n NikeShoe) SayName() string {
  return "Nike Shoe"
}

// Shirt é uma interface para objetos de camisa.
type Shirt interface {
  SayName() string
}

// AdidasShirt é uma implementação concreta de Shirt.
type AdidasShirt struct{}

func (a AdidasShirt) SayName() string {
  return "Adidas Shirt"
}

// NikeShirt é uma implementação concreta de Shirt.
type NikeShirt struct{}

func (n NikeShirt) SayName() string {
  return "Nike Shirt"
}

// AdidasFactory é uma implementação concreta de AbstractFactory que cria objetos da marca Adidas.
type AdidasFactory struct{}

func (a AdidasFactory) CreateShoe() Shoe {
  return AdidasShoe{}
}

func (a AdidasFactory) CreateShirt() Shirt {
  return AdidasShirt{}
}

// NikeFactory é uma implementação concreta de AbstractFactory que cria objetos da marca Nike.
type NikeFactory struct{}

func (n NikeFactory) CreateShoe() Shoe {
  return NikeShoe{}
}

func (n NikeFactory) CreateShirt() Shirt {
  return NikeShirt{}
}

func main() {
  // Obtém uma fábrica para a marca Adidas.
  factory := AdidasFactory{}

  // Cria um sapato e uma camisa da marca Adidas.
  shoe := factory.CreateShoe()
    shirt := factory.CreateShirt()

  // Imprime o nome do sapato e da camisa.
  fmt.Println(shoe.SayName())
  fmt.Println(shirt.SayName())
}
