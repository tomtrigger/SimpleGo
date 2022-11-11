/**
 * iota,特殊常量，可以认为是一个可以被编译器修改的产量
 * iota 在 const 关键字出现时将被重置为0（const 内部的第一行之前），const中每新增一行常量
 * 声明将使 iota 计数一次（iota 可理解为 const 语句块中的行索引）
 * */

package main

import "fmt"

func main() {
	const (
		a = iota	// 0
		b 			// 1
		c 			// 2
		d = "ha"	// 独立值，iota += 1
		e 			// "ha",iota += 1
		f = 100		// iota += 1
		g 			// 100 iota += 1
		h = iota	// 7, 恢复计算
		i 			// 8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		j = 100
		k
		l = iota
	)
	fmt.Println(j, l)

	// 定义常量组时，如果不提供初始值，则表示将使用上行的表达式
	const (
		aa = 100
		bb
	)
	fmt.Println(aa, bb)

	// iota 只在同一个 const 常量组内递增，每当有新的const关键字时，iota计数会重新开始
	const (
		ii = iota
		jj = iota
		kk = iota
	)

	const xx = iota
	const yy = iota

	println(ii, jj, kk, xx, yy)

}