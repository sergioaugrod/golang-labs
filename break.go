package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	counter := 0

	for {
		counter++

		number := rand.Intn(4200)
		fmt.Println(number)

		if number%42 == 0 {
			break
		}
	}

	fmt.Printf("Interactions = %d", counter)
}
