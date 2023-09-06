package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

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

func main() {
	t1 := titulo("https://www.uol.com.br", "https://www.google.com")
	t2 := titulo("https://www.estadao.com.br", "http://g1.globo.com/")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
