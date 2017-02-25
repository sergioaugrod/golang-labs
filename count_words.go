package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		counter, found := statistics[initial]

		if found {
			statistics[initial] = counter + 1
		} else {
			statistics[initial] = 1
		}
	}

	return statistics
}

func PrintCount(statistics map[string]int) {
	fmt.Println("Count of words beginning with the letter:")

	for initial, counter := range statistics {
		fmt.Printf("%s = %d\n", initial, counter)
	}
}

func main() {
	words := os.Args[1:]
	statistics := GenerateStatistics(words)
	PrintCount(statistics)
}
