package main

import (
	"fmt"
	"time"
)

func main() {
	// Declara um canal e um timer
	ch := make(chan int)
	timer := time.NewTimer(time.Second)

	// Inicia uma goroutine para enviar um valor ao canal
	go func() {
		ch <- 1
	}()

	// Usa o `select` para esperar por um valor no canal ou pelo término do timer
	select {
	case v := <-ch:
		fmt.Println("Recebi 1 do canal:", v)
	case <-timer.C:
		fmt.Println("O timer expirou")
	}
}
