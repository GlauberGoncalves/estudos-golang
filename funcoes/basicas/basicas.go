package main

import "fmt"

func funcaoSemParametro() {
	fmt.Println("Função sem parametro")
}

// pode colocar os parametros e só depois por o tipo
func funcaoComParametro(nome, frase string) {
	fmt.Println(nome, frase)
}

func funcaoComRetorno() string {
	return "Retornei algo interessante"
}

func retornaMaisDeUmValor() (string, bool) {
	return "retorno 2 valores pq posso", true
}

func main() {

	funcaoSemParametro()
	funcaoComParametro("Glauber", "você é de mais")
	retorno := funcaoComRetorno()
	retorno1, retorno2 := retornaMaisDeUmValor()

	if retorno2 {
		fmt.Println(retorno1)
		fmt.Println(retorno)
	}

}
