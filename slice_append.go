package main

import "fmt"

func main() {
	numbers := make([]int, 0)
	numbers = append(numbers, 23)
	numbers = append(numbers, 24)

	fmt.Println(numbers)
}
