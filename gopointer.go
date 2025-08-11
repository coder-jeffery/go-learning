package main

import "fmt"

// 结构体指针
type student struct {
	name, address string
	sex           string
	weight        int
}

// 结构体占用一块连续的内存
type she struct {
	s int8
	h int8
	e int8
}

// Go语言指针： &取地址 *根据内存地址取值
func main() {
	//值类型：int float bool string array struct
	//指针类型： *int, *int64, *string

	a := 10
	b := &a
	fmt.Printf("%v  %p \n", a, &a)
	fmt.Println(b)
	c := *b //根据地址取值
	fmt.Printf("%v  %p \n", c, &c)
	fmt.Println(c)

	//
	var x *int
	x = new(int)
	*x = 100
	fmt.Printf("%v  %p \n", x, &x)
	fmt.Println(*x)

	//new make分配内存  new用来分配基本数据类型 整型 浮点型 bool值
	//make slice / map / chan  切片

	//匿名结构体
	var user struct {
		name string
		age  int8
	}

	user.age = 19
	user.name = "jack"
	fmt.Println(user)

	//结构体指针
	var stu = new(student)
	fmt.Printf("%#v \n", stu)
	//(*stu).sex = "male"
	//(*stu).weight = 100
	//(*stu).name = "jack"
	//(*stu).address = "i come from china"

	stu.sex = "male"
	stu.weight = 100
	stu.name = "jack"
	stu.address = "i come from china"

	fmt.Printf("%#v \n", stu)

	//结构体初始化：
	//键值对初始化
	//值列表初始化
	//结构体内存占用一块连续内存

	whoshe := she{1, 2, 3}
	fmt.Printf("%#v \n", whoshe)
	fmt.Printf("%p \n", &whoshe.s)
	fmt.Printf("%p \n", &whoshe.h)
	fmt.Printf("%p \n", &whoshe.e)
}
