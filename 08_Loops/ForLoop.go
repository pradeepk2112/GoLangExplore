package main

import (
	"fmt"
)

func main() {

	streams := []int{4, 3, 6, 8, 2, 5, 11, 7, 5, 3}

	for _, v := range streams {

		if v > 6 {
			fmt.Println("Caught!!")
			break
		} else {
			fmt.Println(v)
		}
	}
}
