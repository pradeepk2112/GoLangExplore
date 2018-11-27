package main

import (
	"fmt"
	"strconv"
)

func main() {

	rect1 := make([]string, 10, 10)

	rect2 := make(map[int]string)

	rect2[1] = "Catalog"
	rect2[2] = "Pricing"
	rect2[3] = "Fulfilment"
	rect2[4] = "Configurator"

	mapper := func(map[int]string) []string {

		for k, v := range rect2 {
			rect1[k] = strconv.Itoa(k) + "--> " + v

		}
		return rect1
	}

	fmt.Printf("%v\n", mapper(rect2))

}
