package main

import "fmt"

type Nota float64

func (n Nota) getConceito() string {
	if n < 6 {
		return "E"
	}
	if n <= 7 {
		return "D"
	}
	if n <= 8 {
		return "C"
	}
	if n <= 9 {
		return "B"
	}
	if n <= 10 {
		return "A"
	}
	return "E"
}

func main() {

	var n Nota
	n = 7.5

	fmt.Println("Nota conceito ", n.getConceito())
}
