package main

import "fmt"

// 数组类型数据  数组长度固定
func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	//定义使用初始值列表
	var cityArray = [4]string{"beijing", "shanghai", "wangwu", "tianjing"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[3])

	//编译器推导
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)
	fmt.Println(len(boolArray))
	fmt.Println(cap(boolArray))

	//索引值初始化
	var titleArray = [...]string{1: "xxx", 2: "yyy", 3: "zzz"}
	fmt.Println(titleArray)

	//数组遍历
	var cityArr = [...]string{"北京", "上海", "深证"}
	for i := 0; i < len(cityArr); i++ {
		fmt.Println(cityArr[i])
	}

	//遍历
	for index, value := range cityArr {
		fmt.Println(index, value)
	}

	//遍历
	for _, value := range titleArray {
		fmt.Println(value)
	}

	//二维数组
	cityArrayList := [3][2]string{
		{"江苏", "连云港"},
		{"浙江", "杭州"},
		{"安徽", "合肥"},
	}
	fmt.Println(cityArrayList[0][1])

	//二维数组遍历
	for _, value := range cityArrayList {
		fmt.Println(value)

		for index, value := range value {
			fmt.Println(index, value)
		}
	}
	//数组属于值类型  数组是值类型
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)

}

func f1(a [3][2]int) {
	a[0][0] = 100
}
