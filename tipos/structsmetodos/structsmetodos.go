package main

import(
	"fmt"
)

type Carro struct {
	marca string
	portas int
	ano int
}

func (p Carro) getMarca() string {
	return p.marca
}

func (p Carro) getQuantidadePortas() int {
	return p.portas
}

func (p Carro) getAno() int {
	return p.ano
}

func main(){

	carro1 := Carro{
		marca: "Honda",
		portas: 4,
		ano: 2019,
	}

	fmt.Println("Marca:", carro1.getMarca())
	fmt.Printf("O carro tem %d portas\n", carro1.getQuantidadePortas())
	fmt.Println("Este carro é do ano de", carro1.getAno())
}