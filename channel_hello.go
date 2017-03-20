package main

import "fmt"

func produce(c chan int) {
	c <- 33
}

func main() {
	c := make(chan int)
	go produce(c)

	value := <-c
	fmt.Println(value)
}
