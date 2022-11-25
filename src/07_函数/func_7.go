package main

import (
	"fmt"
)

// 定义结构体
type Cricle struct {
	radius float64
}

func main() {
	var c Cricle
	fmt.Println(c.radius) // 0

	c.radius = 10.00
	fmt.Println(c.radius) // 10

	c.changeRadius(20)
	fmt.Println(c.radius) // 20

	change(&c, 30)
	fmt.Println(c.radius) // 30
}

func (c Cricle) getArea() float64 {
	return c.radius * c.radius
}

// 如果想要更改成功 c 值，这里需要传指针
func (c *Cricle) changeRadius(radius float64) {
	c.radius = radius
}

// 以下操作将不成功
/*func (c Cricle) changeRadius(radius float64) {
	c.radius = radius
}*/ 

// 引用类型要想改变值需要传指针
func change(c *Cricle, radius float64) {
	c.radius = radius
}
