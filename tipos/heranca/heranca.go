package main

import(
	"fmt"
)

type Pessoa struct {
	pessoaId int
	nome string		
}

type Juridica struct {
	Pessoa
	cnpj string
}

type Fisica struct {
	Pessoa
	cpf string
}

func main() {
	
	pessoa1 := Fisica{}
	pessoa1.pessoaId = 1
	pessoa1.nome = "Glauber"
	pessoa1.cpf = "11832778738"

	pessoa2 := Juridica{}
	pessoa2.pessoaId = 1
	pessoa2.nome = "Google"
	pessoa2.cnpj = "118299920001"
	
	fmt.Println("Id:", pessoa1.pessoaId)
	fmt.Println("Nome:", pessoa1.nome)
	fmt.Println("CPF", pessoa1.cpf)	

	fmt.Println()
	
	fmt.Println("Id:", pessoa2.pessoaId)
	fmt.Println("Nome:", pessoa2.nome)
	fmt.Println("CNPJ", pessoa2.cnpj)
}