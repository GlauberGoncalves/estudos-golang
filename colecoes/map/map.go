package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[1001] = "Glauber"
	aprovados[1002] = "Jessica"
	aprovados[1003] = "Zelia"

	for n, aprovado := range aprovados {
		fmt.Printf("%d, %s\n", n, aprovado)
	}

}
