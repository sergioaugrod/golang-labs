package main

import "fmt"

func sum(n, m int) int {
	return n + m
}

func priceFinal(costPrice float64) (priceDolar float64, priceReal float64) {
	profit := 1.33
	conversionRate := 2.34

	priceDolar = costPrice * profit
	priceReal = priceDolar * conversionRate

	return
}

func main() {
	result := sum(10, 20)
	fmt.Println("The sum is:", result)

	priceDolar, priceReal := priceFinal(34.99)

	fmt.Println("Dolar: ", priceDolar)
	fmt.Println("Real: ", priceReal)
}
