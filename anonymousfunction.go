package main

import "fmt"

// 匿名函数 + 闭包
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x + y
	}
}

// 定义一个函数 返回值是一个函数
func sayHelloBody() func() {
	name := "婆罗门"
	return func() {
		fmt.Println("hello sayHelloBody: ", name)
	}
}

func sayHelloBody2(name string) func() {
	//name := "婆罗门"
	return func() {
		fmt.Println("hello sayHelloBody2: ", name)
	}
}

func main() {
	//闭包 = 函数 + 外层变量引用
	func(x, y int) {
		fmt.Println(x, y)
	}(10, 20)

	var f = adder()
	fmt.Println(f)
	fmt.Println(f(10))
	fmt.Println(f(20))
	s := sayHelloBody()
	s() //执行函数内部的匿名函数
	s2 := sayHelloBody2("this is english name")
	s2()
}
