package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"apple"}
	prices := [4]float64{10.77, 1.99, 2.99, 3.99}
	fmt.Println(prices)
	productNames[2] = "banana"
	fmt.Println(productNames)
	fmt.Println(productNames[0])

	featurdPrices := prices[1:] // slice index 0 to 2 (before 3)
	highlightedPrices := featurdPrices[:1]
	fmt.Println(highlightedPrices)
}
