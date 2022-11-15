package main

import "fmt"

func main() {
	// 定义局部变量
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == "C":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	case grade == "F":
		fmt.Println("不及格")
	default:
		fmt.Println("差")

	}

	fmt.Println("等级是：", grade)

}
