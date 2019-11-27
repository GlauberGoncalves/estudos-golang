package main

import (
	"fmt"
	"time"
)

func multiplica(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go multiplica(2, c)

	a, b := <-c, <-c

	fmt.Println(a, b)
	fmt.Println(<-c)
}
