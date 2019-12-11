package main

import(
	"fmt"
	"time"
)

func rotina(c chan int){
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	fmt.Println("Executou")
	c <- 6
}

func main()  {

	ch := make(chan int, 3)

	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
}