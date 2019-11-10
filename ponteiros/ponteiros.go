package main

import (
	"fmt"
)

func main() {

	ano := 1998

	// criação do ponteiro
	var p *int = nil

	p = &ano
	*p++
	ano = 2019

	fmt.Println(ano, *p)

}
