package main

import("fmt")

func main(){
	p1 := Ponto{3.0, 4.0}
	p2 := Ponto{4.0, 2.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))

}