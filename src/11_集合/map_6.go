package main

import (
	"fmt"
)

func main() {
	personSalay := map[string]int {
		"steve":12000,
		"jamie":15000,
	}

	personSalay["mike"] = 9000
	fmt.Println(personSalay)
	// map[jamie:15000 mike:9000 steve:12000]

	newPersonSalay := personSalay
	newPersonSalay["mike"] = 18000
	fmt.Println(personSalay)
	// map[jamie:15000 mike:18000 steve:12000]
}