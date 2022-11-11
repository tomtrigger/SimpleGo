package main

import (
	"fmt"
)

func main() {
    // %d 表示整型数字，%s 表示字符串
    var stockcode = 123
    var enddata = "2020-12-31"
    var url = "Code=%d&enddata=%s"
    var target_url = fmt.Sprintf(url, stockcode, enddata)
    fmt.Println(target_url)	
}