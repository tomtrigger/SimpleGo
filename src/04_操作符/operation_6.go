/**
 * 其他运算符
 * 
 * 运算符 		描述									实例
 * 
 * &			返回变量存储地址						&a ==> 将给出变量的实际地址
 * *			指针变量								*a ==> 是一个指针变量
 * 
 * */

 package main

 import "fmt"

 func main() {
 	var a int = 4
 	var b int32
 	var c float32
 	var ptr *int

 	fmt.Printf("a 的变量类型为 %T \n",a)
 	fmt.Printf("b 的变量类型为 %T \n",b)
 	fmt.Printf("c 的变量类型为 %T \n", c)

 	/**
 	 * & 和 * 运算符实例
 	 * 
 	 * 指针地址、指针类型和指针取值
 	 * 提示：变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址
 	 * 
 	 * ptr := &v  // v 的类型为 T
 	 * 
 	 * 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为 *T，
 	 * 称作 T 的指针类型，* 代表指针。
 	 * 
 	 * 取地址操作符 & 和取值操作符 * 是一对互补操作符，& 取出地址，*根据地址取出地址指向的值。
 	 * 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
 	 * 对变量进行取地址操作使用 & 操作符，可以获取这个变量的指针变量
 	 * 指针变量的值就是指针地址
 	 * 对指针变量进行取值操作使用 * 操作符，可以获得指针变量指向的原变量的值。
 	 * 
 	 * */
 	ptr = &a 
 	fmt.Printf("a 的值为 %d \n", a)
 	fmt.Printf("*ptr 为 %d \n", *ptr)
 	fmt.Println(ptr)


 	// 使用指针变量与不使用的区别
    var e int = 4
    var ptr1 int

    ptr1 = e
    fmt.Println(ptr1)
    e = 15
    fmt.Println(ptr1)

    var f int = 4
    var prt2 *int

    prt2 = &f
    fmt.Println(*prt2)

    f = 5
    fmt.Println(*prt2)
 }