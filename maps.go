package main

import "fmt"

func main() {
	websites := map[string]string{ // map key string to string value
		"Google":   "www.google.com",
		"Youtube":  "www.youtube.com",
		"Facebook": "www.facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Facebook"]) // get value by key

	websites["Twitter"] = "www.twitter.com" // assign new value
	fmt.Println(websites)

	delete(websites, "Youtube") // delete value
	fmt.Println(websites)
}
