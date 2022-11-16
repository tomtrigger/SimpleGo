package main

import "fmt"

func main() {
	// print9x()
	print9xForGoto()
}

func print9x() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, "\t")
		}
		fmt.Println("")
	}
}

func print9xForGoto() {
	n := 1

LOOP:
	for i := 1; i <= 9; i++ {
		if i <= n {
			fmt.Print(i, "*", n, "=", n*i, "\t")
		} else {
			n++
			fmt.Println("")
			goto LOOP
		}
	}
}
