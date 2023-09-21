package main

import (
	"fmt"
)

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// GetSportsFactory eh uma fabrica
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
