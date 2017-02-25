package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func calculate(base float64, control *sync.WaitGroup) {
	defer control.Done()
	n := 0.0

	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}

	fmt.Println(n)
}

func main() {
	initial := time.Now()
	var control sync.WaitGroup
	control.Add(3)

	go calculate(9.37, &control)
	go calculate(6.94, &control)
	go calculate(42.57, &control)

	control.Wait()
	fmt.Printf("Finishing on %s.\n", time.Since(initial))
}
