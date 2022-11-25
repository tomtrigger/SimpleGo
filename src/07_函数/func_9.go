 package main

 import (
 	"fmt"
 )

 func printParams(args ...int) {
 	for _, n := range args {
 		fmt.Println(n)
 	}
 }

 func main() {
 	a, b, c := 1, 2, 3
 	printParams(a, b, c)
 }