package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	switch a {
	case 100:
		fmt.Println("a 的值为:", a)
		fallthrough
	case 150:
		a = 150
		fmt.Println("a 的值为:", a)
		if b == 200 {
			break
		}
		fallthrough
	case 180:
		a = 180
		fmt.Println("a 的值为：", a)
	default:
		fmt.Println("a 的值为：", a)
	}

}
