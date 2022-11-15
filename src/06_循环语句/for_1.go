package main

import "fmt"

func main() {

	var total int

	for i := 1; i <= 10; i++ {
		total += i
	}

	fmt.Println("1 到 10 的总数为：", total)

}
