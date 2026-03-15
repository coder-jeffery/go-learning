package main

import "fmt"

type dog struct{}

func (d dog) say() {
	fmt.Println("dog hello")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("cat hello")
}

// 定义一个抽象类型 接口
type sayer interface {
	//say()
	say()
}

func play(arg sayer) {
	arg.say() //
}

//接口变量

func main() {
	c1 := cat{}
	play(c1)

	d1 := dog{}
	play(d1)

}
