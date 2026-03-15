package main

import "fmt"

//结构体继承

//type Animal interface {}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s move animal from \n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套
}

func (d *Dog) wang() {
	fmt.Printf("%s wang animal from \n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 10,
		Animal: &Animal{
			name: "Jack",
		},
	}
	d1.wang()
	d1.move()
}
