package main

import "fmt"

func add(x, y int) int { return x + y }

func calce(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	res := calce(100, 200, add)
	fmt.Printf("res:%d\n", res)
}

//匿名函数和闭包
