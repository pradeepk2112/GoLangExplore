package main

import (
	"fmt"
)

func ForRange() {

	services := map[int]string{1: "Catalog", 2: "Fulfillment", 3: "Pricing", 4: "Configurator", 5: "ShoppingCart", 6: "Payment"}

	districts := []string{"Chennai", "Kanchipuram", "Nagarcoil", "Coimbatore", "Tirupur"}

	//Simple print
	fmt.Println("")

	length := len(services)

	//Using for range with key
	fmt.Println("\n\n Using for range with key")
	for k, v := range services {
		fmt.Printf("%d : %s\n", k, v)
	}

	//Using for range without key
	fmt.Println("\n\n Using for range without key")
	for _, v := range services {
		fmt.Printf("%s\n", v)
	}

	//For range for slices
	fmt.Println("\n\n For range for slices")

	for _, v := range districts {
		services[length+1] = v
	}

	//After Updating
	fmt.Println("\n\n After Updating")
	for k, v := range services {
		fmt.Printf("%d : %s\n", k, v)
	}
}
