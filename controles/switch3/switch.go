package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "indefinido"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println("Opa")
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
