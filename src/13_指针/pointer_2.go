package main

import (
	"fmt"
)

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Println("a 变量的地址是：", &a)
	fmt.Println("ip 变量的存储地址：", ip)
	fmt.Println("ip 变量的地址：", &ip)
	fmt.Println("ip 变量的值：", *ip)

	// a 变量的地址是： 0xc0000a6058
	// ip 变量的存储地址： 0xc0000a6058
	// ip 变量的地址： 0xc0000ca018
	// ip 变量的值： 20

}