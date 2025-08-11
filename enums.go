package main

import "fmt"

func main() {
	//iota常量自动生成器
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)
	const d = iota //遇到const 重置0
	fmt.Println(d)

	var bb = true
	fmt.Println(bb)

	f := (1 == 2)
	fmt.Println(f)

	var a1 uint8 = 255
	fmt.Println(a1)

	var as byte = 'a'
	fmt.Println(as)

	var apple = "春敏不觉晓"
	var bana = "处处闻啼鸟"
	fmt.Println(apple)
	fmt.Println(&apple)
	fmt.Println(len(apple)) //计算字符串个数
	fmt.Println(apple + "," + bana)

	var ch byte
	var chStr string
	ch = 'a'
	chStr = "a"
	fmt.Println(ch)
	fmt.Println(chStr)
}
