package main

import (
	"fmt"
)

func main() {

Loop:
	fmt.Print("Enter age: ")

	var age int
	fmt.Scanln(&age)

	if age < 18 {

		fmt.Println("Not eligible")
		goto Loop
	} else {
		fmt.Println("Eligible")
	}

}
