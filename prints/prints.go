package main

import(
	"fmt"
)

func main(){
	pi := 3.1415
	valor := fmt.Sprint(pi)
	var texto string
	texto = "qualquer coisa em texto"

	fmt.Println("O valor de x é", 4)
	fmt.Println("O valor de PI é", valor)
	fmt.Println("O valor de pi é " + valor)
	fmt.Printf("O valor de pi é %s \n", valor)	
	fmt.Println(texto)
}