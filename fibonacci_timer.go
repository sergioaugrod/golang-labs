package main

import (
	"fmt"
	"time"
)

func GenerateFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b
			return a
		}

		for i := 0; i < n; i++ {
			fmt.Println("%d", fib())
		}
	}
}

func Timer(function func()) {
	init := time.Now()
	function()
	fmt.Printf("\nExecution time: %s\n\n", time.Since(init))
}

func main() {
	Timer(GenerateFibonacci(8))
	Timer(GenerateFibonacci(48))
	Timer(GenerateFibonacci(88))
}
