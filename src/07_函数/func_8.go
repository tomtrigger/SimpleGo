package main

import (
	"fmt"
)

// 定义结构体
/**
 * 理解为一个内部类，这个内部类，只定义了一个变量radius，
 * 然后通过 getArea() 的这个形式，给它实现一个方法，这样别的地方
 * 就可以调用这个方法了。
 */
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea()) // 圆的面积 =  314

	var c2 Circle
	c2 = getArea2(c1)
	fmt.Println(c2.radius) // 120

	c3 := getArea3()
	fmt.Println(c3.radius) // 0.999
}

// 该 method 属于 Circle 类型对象中的方法
// 这相当于是给 Circle 类 定义了一个方法
func (c Circle) getArea() float64 {
	// c.raidus 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// 这种是把 Circle 类作为参数传递，并返回 Circle 类对象
func getArea2(c Circle) Circle {
	var temp Circle
	temp.radius = c.radius * 12
	return temp
}

// 这种是返回 Circle 类对象
func getArea3() Circle {
	var temp Circle
	temp.radius = 0.999
	return temp
}