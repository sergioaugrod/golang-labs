package main

import (
	"fmt"
	"os"
	"strconv"
)

func partition(numbers []int, pivot int) (minors []int, largers []int) {
	for _, number := range numbers {
		if number <= pivot {
			minors = append(minors, number)
		} else {
			largers = append(largers, number)
		}
	}

	return minors, largers
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	number := make([]int, len(numbers))
	copy(number, numbers)

	indexPivot := len(numbers) / 2
	pivot := number[indexPivot]
	number = append(number[:indexPivot], number[indexPivot+1:]...)

	minors, largers := partition(number, pivot)

	return append(
		append(quicksort(minors), pivot),
		quicksort(largers)...)
}

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, v := range input {
		number, err := strconv.Atoi(v)

		if err != nil {
			fmt.Printf("%s is not a valid number!\n", v)
			os.Exit(1)
		}

		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}
