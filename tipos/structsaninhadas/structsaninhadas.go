package main

import (
	"fmt"
)

type item struct {
	produtoId  int
	quantidade int
	preco      float64
}

type pedido struct {
	usuarioId int
	itens     []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		usuarioId: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.99},
		},
	}

	fmt.Println("O valor total do pedido é", pedido.valorTotal())
}
