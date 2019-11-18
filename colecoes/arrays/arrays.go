package main

import "fmt"

func main() {

	// arrays são estaticos e do mesmo tipo
	var notas [4]float64
	notas[0] = 10
	notas[1], notas[2], notas[3] = 9.5, 8.4, 10.0

	total := 0.0
	var media float64 = 0.0

	for i := 0; i < len(notas); i++ {
		total = total + notas[i]
	}

	media = total / 4.0

	fmt.Println("A media é ", media)

}
