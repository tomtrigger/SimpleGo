package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// go 中格式化字符串并赋值给新串，使用 fmt.Sprintf
	// %s 表示字符串
	var stockcode = "000987"
	var enddata = "2020-12-31"
	var url = "Code=%s&enddata=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddata)
	fmt.Println(target_url)

	// 另外一个实例，%d 表示整型
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s)
}