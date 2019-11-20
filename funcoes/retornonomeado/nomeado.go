package main

import (
	"fmt"
)

func getNome() (nome, sobrenome string) {
	nome = "Glauber"
	sobrenome = "Gonçalves"
	return
}

func main() {
	nome, sobrenome := getNome()

	fmt.Println(nome, sobrenome)
}
