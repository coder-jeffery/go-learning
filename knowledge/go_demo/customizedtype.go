package main

import "fmt"

// 自定义类型
type MyInt int

// 类型别名
type NewInt = int

func main() {
	var i MyInt
	fmt.Printf("type %T  value: %v \n", i, i)
	fmt.Printf("testing...............")

	var j NewInt
	fmt.Printf("type %T  value: %v \n", j, j)
}
