package main  // 主包，程序入口必须在main包

import (
	"fmt"       // 格式化输入输出
	"pet-system/pkg/utils"  // 导入自定义包（基于模块名）
	"time"
)

//闭包 内部函数作为返回值/参数传递
//内部函数引用外部变量
//外部变量周期延长 闭包绑定
func adder() func(int )int{
	sum := 0
	return func(x int )int{
		sum += x
		return sum
	}
}


func main() {
	fmt.Println("Hello, Go Project!")
	// 调用自定义包的函数
	fmt.Println("计算结果：", utils.Add(10, 20))

	fmt.Println(time.DateTime) // 2006-01-02 15:04:05
	fmt.Println(time.Now()) // 2026-03-15 14:58:38.021581 +0800 CST m=+0.000108460


	now := time.Now();
	fmt.Println(now) // 2026-03-15 14:59:38.838316 +0800 CST m=+0.000236585

	//时间类型转换区别
	fmt.Println("标准时间格式， ", now.Format("2006-01-02 15:04:05"))
	fmt.Print("test interested") // print不带换行操作
	fmt.Println("test metadata") // println带换行操作


	//闭包的代码演示
	add  := adder()
	fmt.Println("打印输出的数值：",add(10))
	fmt.Println("打印第二个数值:", add(99))

	add2 := adder()
	fmt.Println("打印输出的数值：",add2(1))
	fmt.Println("打印第二个数值:", add2(99))
}