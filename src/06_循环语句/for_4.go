package main

import "fmt"

func main() {
	strings := []string{"huang", "xiu", "nian"}

	for i, s := range strings {
		fmt.Println(i, s)
	}

	nums := [6]int{1, 2, 3, 4}

	for key, num := range nums {
		fmt.Println(key, num)
	}
}
