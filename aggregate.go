package main

import "fmt"

type Aggregator func(n, m int) int

func Aggregate(values []int, initial int, fn Aggregator) int {
	aggregated := initial

	for _, v := range values {
		aggregated = fn(v, aggregated)
	}

	return aggregated
}

func CalculateSum(values []int) int {
	sum := func(n, m int) int {
		return n + m
	}

	return Aggregate(values, 1, sum)
}

func CalculateMultiplication(values []int) int {
	multiplication := func(n, m int) int {
		return n * m
	}

	return Aggregate(values, 1, multiplication)
}

func main() {
	values := []int{3, -2, 5, 7, 8, 22, 32, -1}

	fmt.Println(CalculateSum(values))
	fmt.Println(CalculateMultiplication(values))
}
