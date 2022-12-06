package main

import (
	"fmt"
)

func main() {
	animals :=[][]string{}

	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	for i := range animals {
		fmt.Printf("Row:%v ", i)
		fmt.Println(animals[i])
	}

}