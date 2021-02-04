package main

import "fmt"

func main() {
	/*
		iota:特殊的常量，可以被编译器自动修改的常量
			每当定义一个const，iota的初始值为0
			每当定义一个常量，就会自动累加1
			直到下一个const出现，清零
	*/
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const (
		d = iota // 0
		e
	)
	fmt.Println(d)
	fmt.Println(e)

	// 枚举中
	const (
		MALE   = iota //0
		FEMALE        // 1
		UNKNOW        // 2
	)
	fmt.Println(MALE, FEMALE, UNKNOW)

	fmt.Println("------------------")
	const (
		A = iota   // 0
		B          // 1
		C          // 2
		D = "HAHA" // iota = 3
		E          // haha iota = 4
		F = 100    // iota = 5
		G          // 100 iota = 6
		H = iota   // 7
		I          // iota 8
	)
	const J = iota // 0

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)

}
