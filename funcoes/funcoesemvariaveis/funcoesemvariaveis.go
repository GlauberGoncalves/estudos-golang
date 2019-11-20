package main

import (
	"fmt"
)

var somaDoisNumeros = func(a, b int) (c int) {
	c = a + b
	return
}

func main() {
	resultado := somaDoisNumeros(10, 14)
	fmt.Println(resultado)

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(10, 4))

}
