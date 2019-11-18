package main

import "fmt"

func main() {

	array := [...]int{1, 2, 3, 4, 5, 6, 7}
	slice := array[2:5] // referencia ao pedaço do array original

	for _, numero := range slice {
		fmt.Println(numero)
	}
	fmt.Println("\n")

	slice2 := slice[0:1]

	for _, numero := range slice2 {
		fmt.Println(numero)
	}
}
