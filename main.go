package main

import "fmt"

func main() {
	userNames := make([]string, 2) // make and array of length 2
	userNames = append(userNames, "hello")

	userNames[0] = "satu"
	userNames[1] = "dua"
	// userNames[2] = "tiga"

	fmt.Println(userNames)
}
