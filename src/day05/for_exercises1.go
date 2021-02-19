package main

import "fmt"

func main() {
	/*
		for循环的练习题：
		练习1：打印58-23数字
		练习2：求1-100的和
			sum := 0
			sum = +1+2+3...+100
			sum = sum + 1
		练习3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
	*/

	for i := 58; i >= 28; i-- {
		fmt.Print(i, ",")
	}

	fmt.Println("----------------------")

	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和", sum)

	fmt.Println("-----------------------")

	count := 0 // 计数器
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print("\t", i)
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println("count--->", count)

}
