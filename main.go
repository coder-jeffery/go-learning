package main

import (
	"fmt"
)

// 命名规范 ：  1.数字 字母 下划线  2.不允许用关键字 3.区分大小写  4. 不能使用%@ 不可以数字开头
func main() {
	//变量声明必须使用
	var num int
	var name string
	fmt.Println(num)
	fmt.Println(name)

	num = 10
	fmt.Println(num)

	var n1 float32 = 3.14
	var sum float32 = n1 * n1
	fmt.Println(sum)
	fmt.Printf("打印num数据的类型 %T", sum)

	new_num := 10 //自动推导 左边必须是未声明的变量 初始化只可以一次
	fmt.Println("\n", new_num)
	fmt.Println("hello world")

	a, b, c := 10, 9.99, true
	fmt.Println("\n", a, b, c)
	fmt.Printf("自动推导类型 %T %T %T", a, b, c)
}
