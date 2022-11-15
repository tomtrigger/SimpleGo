package main

import "fmt"

func main() {
	sum := 1

	for ; sum <= 10; {
		sum += sum
	}

	fmt.Println(sum)

	sum = 1

	// 可以像下面这样写，更像 while 语句
	for sum <= 10 {
		sum += sum
	}

	fmt.Println(sum	)
}
