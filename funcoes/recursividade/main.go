package main

import (
	"fmt"
)

func fibonat(n1 int, n2 int, quantidade int) error {

	if quantidade == 0 {
		return fmt.Errorf("Quantidade 0")
	}

	if n1 > n2 {
		return fmt.Errorf("Primeiro numero é maior que segundo")
	}

	fmt.Printf("%d\n", n1+n2)
	fibonat(n2, n1+n2, quantidade-1)
	return nil
}

func main() {
	fibonat(0, 1, 50)
}
