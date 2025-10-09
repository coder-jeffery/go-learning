package main

import "fmt"

// 函数进阶之变量作用域

// 定义全局变量
var num int = 10

// 定义函数
func testGlobal() {
	num := 100
	fmt.Println("Global", num)

	//全局变量
	//局部变量
}

func main() {
	//
	testGlobal()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc() //函数可以作为变量

}
