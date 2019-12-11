package main

import(
	"fmt"
	"time"
)

func rotina(c chan int){
	time.Sleep(time.Second)
	fmt.Println("so depois foi lido")
}

func main(){
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")

}