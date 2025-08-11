package main

import (
	"fmt"
	"os"
)

func showMenu() {

	fmt.Println("------------------------")
	fmt.Println("管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.展示系统")
	fmt.Println("4.退出系统")

}

func main() {
	for {
		showMenu()
		var input int
		fmt.Print("请输入：")
		fmt.Scanf("%d \n", &input)
		fmt.Println("您输入的数字:", input)
		switch input {
		case 1:
		//
		case 2:
		//
		case 3:
		//
		case 4:
			os.Exit(0)

		}
	}
}
