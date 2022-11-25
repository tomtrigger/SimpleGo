package main

import (
	"fmt"
)

func main() {
	// 常量用作枚举
	const (
		UNKNOWN = 0
		FEMALE  = 1
		MALE    = 2
	)

	var gender = UNKNOWN

	switch gender {
	case UNKNOWN:
		fmt.Println("性别是：未知")
	case FEMALE:
		fmt.Println("性别是：女性")
	case MALE:
		fmt.Println("性别是：男性")
	default:
		fmt.Println("没有性别")
		return
	}

}
