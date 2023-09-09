package main

import (
	"fmt"
	"time"

	"github.com/msn2edson/html"
)

func oMaisRapido(url1, url2, url3 string) string {

	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		fmt.Println("t1", t1)
		return t1
	case t2 := <-c2:
		fmt.Println("t2", t2)
		return t2
	case t3 := <-c3:
		fmt.Println("t3", t3)
		return t3
	case <-time.After(2000 * time.Millisecond): // timeout de 5 segs
		return "Todos perderam"
		//default:
		//		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.uol.com",
	)
	fmt.Println(campeao)
}
