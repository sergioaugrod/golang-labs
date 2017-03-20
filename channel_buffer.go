package main

import "fmt"

func produce(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func main() {
	c := make(chan int, 3)
	go produce(c)

	for value := range c {
		fmt.Println(value)
	}
}
