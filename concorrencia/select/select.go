package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

// Isso eh um generator
// <-chan - canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) // Importante: foi criado a funcao e agora estou invocando
	}
	return c
}

func oMaisRapido(url1, url2, url3 string) string {

	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		fmt.Println("t1 " + t1)
		return t1
	case t2 := <-c2:
		fmt.Println("t2 " + t2)
		return t2
	case t3 := <-c3:
		fmt.Println("t3 " + t3)
		return t3
	case <-time.After(5000 * time.Millisecond): // timeout de 5 segs
		return "Todos perderam"
		//default:
		//		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.facebook.com",
	)
	fmt.Println(campeao)
}
