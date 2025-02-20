package main

import "fmt"

func main() {
	// make in slices
	userNames := make([]string, 2, 5) // make an array of length 2 and capacity 5
	userNames = append(userNames, "hello")

	userNames[0] = "satu"
	userNames[1] = "dua"
	// userNames[2] = "tiga"

	fmt.Println(userNames)
	fmt.Println(cap(userNames))

	courseRatings := make(map[string]float64, 3) // make a map of length 3 (use make if you already know the length you will use)
	courseRatings["Go"] = 4.5
	courseRatings["React"] = 5
	courseRatings["Angular"] = 4
	fmt.Println(courseRatings)
}
