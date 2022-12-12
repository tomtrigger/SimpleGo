package main

import (
	"fmt"
)

func main() {
	countryCapitalMap := map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

	fmt.Println("删除前")
	for contry := range countryCapitalMap {
		fmt.Println(contry, countryCapitalMap[contry])
	}
	// France Paris
	// Italy Rome
	// Japan Tokyo
	// India New Delhi

	delete(countryCapitalMap, "France")

	fmt.Println("删除后")
	for contry := range countryCapitalMap {
		fmt.Println(contry, countryCapitalMap[contry])
	}
	// Italy Rome
	// Japan Tokyo
	// India New Delhi

}