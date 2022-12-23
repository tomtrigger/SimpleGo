package main

import (
	"fmt"
)

func main() {
	// 使用 range 去求一个 slice 的和.使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 在数组上使用 range 将传入索引和值两个变量
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 在 map 数据类型上的使用
	kvs := map[string]string{"a": "apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	// range 也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}