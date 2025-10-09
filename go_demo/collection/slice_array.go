package main

import (
	"fmt"
	"sort"
)

// 切片  对底层数组的封装
func main() {

	//基于数组切片
	var a []string
	var b []int
	fmt.Println("a:", a, "b:", b)

	//切片之后再切片
	a1 := [5]int{55, 56, 57, 58, 59}
	b1 := a1[1:4] //切片
	fmt.Println(b1)
	fmt.Println("b1:", b1)
	fmt.Printf("%T \n", b1)

	c1 := b1[1:3]
	fmt.Println(c1)

	//make函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)

	//切片的赋值拷贝
	s1 := make([]int, 3)
	fmt.Println(s1)
	s2 := s1
	s2[0] = 100
	fmt.Println(s2)
	fmt.Println(len(s1) == 0)
	//一个nil值的切片没有底层数组  一个nil值切片长度和容量是0

	//切片扩容
	var ex []int //没有申请内存
	ex = append(ex, 10)
	fmt.Println(ex)
	fmt.Println("*************************")

	var ex2 []int
	for i := 0; i < 10; i++ {
		ex2 = append(ex2, i)
		fmt.Println("%v  len:%d  cap:%d  ptr:%p \n", ex2, len(ex2), cap(ex2), ex2)
	}

	fmt.Println("*************************")

	//切片扩容
	//切片copy

	s := []int{1, 2, 3, 4, 5}
	m := make([]int, 5, 5)
	l := m
	copy(m, s)
	m[0] = 100
	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(l)

	//切片元素删除
	aa := []string{"北京", "上海", "深圳", "广州"}
	aa = append(aa[0:2], aa[3:]...)
	fmt.Println(aa)

	//切片 内置sort排序
	var bb = [...]int{1, 3, 9, 6, 7, 5}
	sort.Ints(bb[:])
	fmt.Println(bb)
	sort.Reverse(sort.IntSlice(bb[:]))
	fmt.Println(bb)
}
