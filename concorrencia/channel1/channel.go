package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1           // enviando dados para o canal
	fmt.Println(<-ch) // lendo dados do canal

	ch <- 2           // enviando dados para o canal
	fmt.Println(<-ch) // lendo dados do canal

}
