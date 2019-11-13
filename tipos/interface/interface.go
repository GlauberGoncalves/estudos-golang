package main

import (
	"fmt"
)

type imprimivel interface {
	toString() string
}

type carro struct {
	marca string
	valor float64
}

type moto struct {
	marca string
	valor float64
}

func (c carro) toString() string {
	return fmt.Sprintf("%s, %f", c.marca, c.valor)
}

func (m moto) toString() string {
	return fmt.Sprintf("%s, %f", m.marca, m.valor)
}

func imprimir(obj imprimivel) {
	fmt.Println(obj.toString())
}

func main() {

	c := carro{"ferrari", 9999999}
	m := carro{"honda", 30000}
	
	imprimir(c)
	imprimir(m)
}
