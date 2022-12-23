package main 

import (
	"time"
)

// 定义 time.Duration 的别名为 MyDuration
type MyDuration = time.Duration

// 为 MyDuration 添加一个函数
func (m MyDuration) EasySet(a string) { // cannot define new methods on non-local type time.Duration

}

func main() {

}

