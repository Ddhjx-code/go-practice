package main

import (
	"fmt"
)

func main() {
	var a = "test-init" //string,字符串

	var b, c int = 0, 5 //int,整型

	var d = true //bool 布尔型

	var e float64 //float 浮点类型

	//template variables 临时变量
	f := float32(e)

	fmt.Printf(a, b, c, d, e, f)
}
