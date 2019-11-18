package main

import "fmt"

func main() {

	inteiros := [...]int{9, 8, 7, 6, 5} // compilador faz conta

	for i, inteiro := range inteiros {
		fmt.Println(i, " inteiro é ", inteiro)
	}

}
