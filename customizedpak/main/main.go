package main

import (
	"fmt"
	"go_learning/customizedpak/calc"
)

// _ 匿名导入包  init初始化函数

var test = "test var"

// init 函数 日志初始化  数据加载 | 循环引用问题
// Go语言不允许循环使用  定义包实现代码复用
func init() {
	fmt.Println("init call")
	fmt.Println(test)
}

func main() {
	res := calc.Add(1, 2)
	fmt.Println(res)
}
