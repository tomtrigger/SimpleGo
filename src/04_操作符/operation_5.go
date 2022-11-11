/**
 * 赋值运算符
 * 
 * 运算符 		描述									实例
 * 
 * =			简单的赋值运算符，将一个表达式的值给	C = A + B ==> A + B 表达式结果赋值给 C
 * 				一个左值
 * +=			相加后在赋值							C += A  ==>  C = C + A
 * -=			相减后在赋值							C -= A  ==>	 C = C - A
 * *=			相乘后在赋值							C *= A  ==>  C = C * A
 * /=			相除后在赋值							C /= A  ==>  C = C / A
 * %=			求余后在赋值							C %= A  ==>  C = C % A
 * <<=			左移后在赋值							C <<= 2  ==>  C = C << 2
 * >>=			右移后在赋值							C >>= 2  ==>  C = C >> 2
 * &=			按位与后赋值							C &= 2  ==>  C = C & 2
 * ^=			按位异或后赋值						C ^= 2  ==>  C = C ^ 2
 * |=			按位或后赋值							C |= 2  ==>  C = C | 2	
 * 
 * */

 package main

 import "fmt"

 func main() {
 	var a int = 21
 	var c int

	c = a
 	fmt.Println("c = a:", c)

 	c += a
 	fmt.Println("c += a:", c)

 	c -= a
 	fmt.Println("c -= a:", c)

 	c *= a
 	fmt.Println("c *= a:", c)

 	c /= a
 	fmt.Println("c /= a:", c)

 	c = 200

 	c <<= 2
 	fmt.Println("c <<= 2:", c)

 	c >>= 2
 	fmt.Println("c >>= 2:", c)

 	c &= 2
 	fmt.Println("c &= 2:", c)

 	c ^= 2
 	fmt.Println("c ^= 2:", c)

 	c |= 2
 	fmt.Println("c |= 2:", c)

 }