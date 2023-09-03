package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1999.23, []string{"promocao", "eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Notebook2","preco":1999.23,"tags":["promocao2","eletronico2"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])

}
