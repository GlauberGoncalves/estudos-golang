package main

import (
	"fmt"
)

func comprimenta(nome string) {
	fmt.Printf("Olaaa %s!\n", nome)
}

func dizIdade(idade int) {
	fmt.Printf("Você tem %d anos de idade \n", idade)
}

func calculaIdadePorAnoNascimento(anoNascimento int, anoAtual int) int {
	return anoAtual - anoNascimento
}