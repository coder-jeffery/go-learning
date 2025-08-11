package main

import "fmt"

// 实现构造函数
type person struct {
	name, address string
	age           int
}

func newPerson(name, address string, age int) *person {
	return &person{
		name:    name,
		address: address,
		age:     age,
	}
}

// constructor function
func main() {
	p1 := newPerson("John", "Doe", 18)
	fmt.Printf("p1: %#v\n", p1)
}

//方法和函数区别：
