package main

import (
    "fmt"
)

// 声明全局变量
var g int

func main() {
    // 声明局部变量
    var a, b int

    // 初始化
    a = 10
    b = 20
    g = a + b

    fmt.Println("a + b = ", g)
}