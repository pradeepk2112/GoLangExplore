package main

import (
	"fmt"
)

func main() {

	//Maps are unordered. When retrieving, it is displayed in random order
	streams := make(map[int]string)

	rates := map[int]string{}

	streams[1] = "Java"
	streams[2] = "C++"
	streams[3] = "Python"
	streams[4] = "GoLang"

	rates[1] = "Postman"
	rates[2] = "Fiddler"
	rates[3] = "SOAPUI"
	rates[4] = "SQMS"

	// fmt.Println("Streams")
	// fmt.Println(streams)
	// fmt.Println("Rates")
	// fmt.Println(rates)

	// delete(rates, 1)
	// fmt.Println(rates)

	fmt.Println("Unordered")
	for k, v := range streams {

		fmt.Printf("%d  ---->  %s\n", k, v)
	}

	fmt.Println("Ordered")

	for i := 1; i <= len(streams); i++ {

		fmt.Printf("%d  ---->  %s\n", i, streams[i])
	}
}
