package main

import "fmt"

func fazerBolo() string {
	ingredientes := separarIngredientes()
	massa := misturarIngredientes(ingredientes)
	return assarNoFormo(massa)
}

func separarIngredientes() string {
	fmt.Println("Separando ingredientes")
	return "separados"
}

func misturarIngredientes(ingredientes string) string {
	fmt.Println("Misturando ingredientes")
	return "misturados"
}

func assarNoFormo(massa string) string {
	fmt.Println("Assando a massa")
	return "Bolo pronto"
}

func main() {
	bolo := fazerBolo()
	fmt.Println(bolo)
}
