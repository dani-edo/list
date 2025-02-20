package list

import "fmt"

func main() {
	prices := []float64{10.77, 1.99}
	fmt.Println(prices[0:1])
	prices[1] = 2.

	prices = append(prices, 3.99) // append value in the array,then replace prices value with "prices ="
	prices = prices[1:]           // remove first element
	fmt.Println(prices)

	discountPrices := []float64{101.77, 1.99, 2.99, 3.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"apple"}
// 	prices := [4]float64{10.77, 1.99, 2.99, 3.99}
// 	fmt.Println(prices)
// 	productNames[2] = "banana"
// 	fmt.Println(productNames)
// 	fmt.Println(productNames[0])

// 	featurdPrices := prices[1:] // slice index 1 to the latest index
// 	featurdPrices[1] = 2.00
// 	highlightedPrices := featurdPrices[:1] // slice is just window of original array (still mutable)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(featurdPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // length and capacity

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
