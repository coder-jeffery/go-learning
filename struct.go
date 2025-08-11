package main

import "fmt"

// 定义结构体
type employee struct {
	name string
	age  int
	addr string
}

func main() {
	var j employee
	j.name = "Jack"
	j.age = 23
	j.addr = "Alex"
	fmt.Println("j=%#v \n", j)
}
