package main

import(
	"fmt"
	"math"
)

func main(){	
	areaCircunferencia()
}

func areaCircunferencia(){
	const(
		PI float64 = 3.1415
	)

	var(
		raio = 3.8
	) 

	var nao, sim bool = false, true
	
	area := PI * math.Pow(raio, 2)	
	fmt.Println("A área da circunferência", area)

	fmt.Println(nao, sim)
}