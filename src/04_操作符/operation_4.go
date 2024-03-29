/**
 * 位运算符(&、|、^)
 * 
 * 位运算符对整数在内存中的二进制位进行操作
 * 
 * 
 * p 	q		p&q		p|q		p^q
 * 
 * 0	0		0		0		0
 * 
 * 0	1 		0 		1 		1
 * 
 * 1 	1 		1 		1 		0
 * 
 * 1 	0 		0		1 		1
 * 
 * 
 * 假定 A = 60，B = 13，其二进制数转换为：
 * A = 0011 1100
 * B = 0000 1101
 * 
 * A&B = 0000 1100
 * A|B = 0011 1101
 * A^B = 0011 0001
 * 
 * */

 /**
  * Go 语言支持的位运算符如下表所示。假定 A=60，B=13
  * 
  * 运算符 	描述											实例
  * 
  * &		按位与运算符"&"是双目运算符。其功能是参与		（A&B）结果为12，二进制为0000 1100
  * 		运算的两数各对应的二进制想与。
  * 
  * |		按位或运算符"|"是双目运算符。其功能是参与		（A|B）结果为61，二进制为0011 1101
  * 		运算的两数各对应的二进位相或。
  * 
  * ^		按位异或运算符"^"是双目运算符。其功能是参与		（A^B）结果为49，二进制为0011 0001
  * 		与运算的两数各对应的二进制位相异或，当两对
  * 		应的二进制位相异时，结果为1。
  * 
  * <<		左移运算符"<<"是双目运算符。左移n位就是乘以		 A << 2 结果为240，二进制为1111 0000
  * 		2的n次方。其功能把"<<"左边的运算数的各二进位
  * 		全部左移若干位，由"<<"右边的数指定移动的位数，
  * 		高位丢弃，低位补0。
  * 
  * >> 		右移运算符">>"是双目运算符。右移n位就是除以		 A >> 2 结果为15，二进制为0000 1111
  * 		2的n次方。其功能是把">>"左边的运算数的各二进位
  * 		全部右移若干位，">>"右边的数指定移动的位数。
  * */

  package main

  import "fmt"

  func main() {
  	var a uint = 60  // 60 = 0011 1100
  	var b uint = 13  // 13 = 0000 1101

  	fmt.Println("a & b:", (a & b)) // 0000 1100

  	fmt.Println("a | b:", (a | b)) // 0011 1101

  	fmt.Println("a ^ b:", (a ^ b)) // 0011 0001

  	fmt.Println("a << 2:", (a << 2)) // 1111 0000

  	fmt.Println("a >> 2:", (a >> 2)) // 0000 1111
  }