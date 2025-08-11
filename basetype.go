package main

import "fmt"

func test() {

	//声明变量
	var name string
	var age int

	//批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(&name, &age)
	fmt.Println(&a, &b, &c, &d)

	//声明同时赋值
	var ss string = "hello world"
	fmt.Println(ss)

	//类型推导
	var str_name = "jeffery"
	var str_age = 20
	fmt.Println(str_name, str_age)

	//简短变量声明
	m := 10
	fmt.Println(m)

}
