package main

import (
	"fmt"
)

func main() {

	lista1 := [3]int{1, 2, 3}
	var lista2 []int

	lista2 = append(lista2, 8, 7, 6)
	lista2 = append(lista2, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	fmt.Println(lista1, lista2)

	var lista3 []int

	copy(lista3, lista2)

	fmt.Println(lista3)

}
