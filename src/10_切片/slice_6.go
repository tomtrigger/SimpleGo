package main

import (
	"fmt"
)

func main() {
	numa := [3]int{78, 79, 80}

	// 创建切片
	nums1 := numa[:]
	nums2 := numa[:]

	fmt.Println("array before change 1", numa)
	// array before change 1 [78 79 80]

	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	// array after modification to slice nums1 [100 79 80]

	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	// array after modification to slice nums2 [100 101 80]	
}