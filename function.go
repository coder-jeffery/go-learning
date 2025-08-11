package main

import "fmt"

// function函数的使用

func sayHello(name string) {
	fmt.Println("hello this is first name : Robin")
}

func sayHello2(name string) {
	fmt.Println("Hello2 this is second name: ", name)
}

// Go语言参数简写
func calcSum(i, j int) int {
	return i + j
}

// 定义多参数返回值函数 多返回值简写
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 可变参数
func calcSumChange(k ...int) int {
	ret := 0
	for _, arg := range k {
		ret += arg
	}
	return ret
}

// 固定参数 +  可变参数
func calcSumChange2(l int, k ...int) int {
	ret := l
	for _, arg := range k {
		ret += arg
	}
	return ret
}

//固定参数和可变参数同时出现  可变参数放在最后

func main() {
	n := "Jeffery"
	sayHello(n)
	sayHello2("Jeffery")
	c := calcSum(1, 2)
	fmt.Println(c)

	//可变参数
	k1 := calcSumChange(10)
	k2 := calcSumChange(10, 20)
	k3 := calcSumChange(10, 20, 30)
	fmt.Println(k1, k2, k3)

	fmt.Println("******************function test***********************")
	kk1 := calcSumChange2(1, 2, 3, 4, 5)
	kk2 := calcSumChange2(5, 10, 20, 30)
	fmt.Println(kk1)
	fmt.Println(kk2)

	fmt.Println("******************function test***********************")

	x1, x2 := calc(50, 40)
	fmt.Println(x1, x2)
}
