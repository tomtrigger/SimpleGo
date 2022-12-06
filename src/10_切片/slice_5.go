package main

import (
	"fmt"
)

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 59}
	dslice := darr[2:5]
	
	fmt.Println("array before:", darr)
	for key, value := range dslice {
		value++
		dslice[key] = value
	}
	fmt.Println("array after:", darr)

	// array before: [57 89 90 82 100 78 67 59]
	// array after: [57 89 91 83 101 78 67 59] 
}