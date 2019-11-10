package main

import(
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	sexo string	
}

func (p pessoa) getNome() string{
	return p.nome
}

func main() {
	var pessoa1 pessoa
	pessoa1 = pessoa{
		nome: "glauber",
		sobrenome: "gonçalves",
		sexo: "masculino",
	}

	fmt.Println(pessoa1.getNome())
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa1.sexo)

}