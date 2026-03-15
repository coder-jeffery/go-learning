package main

import "fmt"

type Son struct {
	name string
	age  int8
}

func newSon(name string, age int8) *Son {
	return &Son{
		name: name,
		age:  age,
	}
}

func (s Son) BuildDream() {
	fmt.Printf("%s dream: change history \n", s.name)
}

func main() {
	s1 := newSon("jack", int8(18))
	s1.BuildDream()
}

//函数传参 值拷贝
//方法的定义 需要额外添加一个接收者 ： 值接收者 指针接收者 （大对象拷贝/ 如果定义若干个方法 有一个方法是 其他最好也是）
