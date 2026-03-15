package main

import "fmt"

func main() {

	//写法一
	var socre int = 65
	if socre > 90 {
		fmt.Println("A")
	} else if socre > 70 && socre < 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	fmt.Println("***************************************")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("***************************************")

	sv := 1
	switch sv {
	case 1, 3, 5, 7, 9:
		fmt.Println("1")
	case 2, 4, 6, 8, 10:
		fmt.Println("2")
	default:
		fmt.Println("无效输入")
	}

	fmt.Println("***************************************")
	age := 19
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age > 18:
		fmt.Println("大于18 成年了")
	}

}
