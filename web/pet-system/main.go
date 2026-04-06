package main  // 主包，程序入口必须在main包

import (
	"fmt" // 格式化输入输出
	"pet-system/internal/model"
	"pet-system/pkg/utils" // 导入自定义包（基于模块名）
	"sync"
	"time"
)

//闭包 内部函数作为返回值/参数传递
//内部函数引用外部变量
//外部变量周期延长 闭包绑定
func adder() func(int ) int{
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func Divide(a, b int) (int,  error){
	if b ==0 {
		return 0, fmt.Errorf("xxx b  = %b", b)
	}

	return a / b , nil
}


func main() {
	fmt.Println("Hello, Go Project!")
	// 调用自定义包的函数
	fmt.Println("计算结果：", utils.Add(10, 20))

	fmt.Println(time.DateTime) // 2006-01-02 15:04:05
	fmt.Println(time.Now())    // 2026-03-15 14:58:38.021581 +0800 CST m=+0.000108460

	now := time.Now();
	fmt.Println(now) // 2026-03-15 14:59:38.838316 +0800 CST m=+0.000236585

	//时间类型转换区别
	fmt.Println("标准时间格式， ", now.Format("2006-01-02 15:04:05"))
	fmt.Print("test interested") // print不带换行操作
	fmt.Println("test metadata") // println带换行操作

	//闭包的代码演示
	add := adder()
	fmt.Println("打印输出的数值：", add(10))
	fmt.Println("打印第二个数值:", add(99))

	add2 := adder()
	fmt.Println("打印输出的数值：", add2(1))
	fmt.Println("打印第二个数值:", add2(99))

	add3 := adder()
	fmt.Println("打印输出的数值：", add3(1))
	fmt.Println("打印输出的数值：", add3(100))

	p := model.Pet{
		ID:     99999,
		Name:   "Tim",
		Weight: 0,
		Color:  "",
		Birth:  time.Location{},
		Dead:   false,
	}
	fmt.Printf("id: %d , name: %s", p.ID, p.Name)

	fmt.Println("\n")

	p2 := model.Pet{ID: 88888, Name: "jim", Weight: 11.1, Color: "red", Birth: time.Location{}, Dead: false}
	fmt.Printf("id : %d, name : %s", p2.ID, p2.Name)

	fmt.Println(p.Getinfo())
	//fmt.Println(p2.Getinfo())
	fmt.Println(p2.Update("this is update my name: Jeffery"))
	//fmt.Println(p2.Update(""))
	//
	//
	b1 := model.Bird{Name: "小鸟"}
	p1 := model.Plane{Model: "波音131"}

	model.LetFly(p1)
	model.LetFly(b1)

	dog1 := model.Dog{Name : "bard"}
	model.CommonPet(dog1)

	//error

	result, err  := Divide(2, 1)
	if err != nil{
		fmt.Println("错误信息", err)
		return
	}
	fmt.Println("结果: >>>>>", result)


///协程
	go func() {

		for i := 0; i< 5; i++ {
			fmt.Printf("协程执行 %d \n ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("主协程执行")
	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")


	//
	// 定义等待组，用于等待所有协程执行完成
	var wg sync.WaitGroup
	// 目标：启动100万个协程
	total := 1000000
	wg.Add(total)

	// 记录启动开始时间
	start := time.Now()

	// 循环启动协程
	for i := 0; i < total; i++ {
		go func(num int) {
			defer wg.Done() // 协程执行完，等待组计数-1
			// 模拟简单计算（无IO，纯CPU操作）
			_ = num * num
		}(i)
	}

	// 等待所有协程执行完成
	wg.Wait()

	// 计算耗时
	elapsed := time.Since(start)
	fmt.Printf("启动%d个协程，总耗时：%v\n", total, elapsed)
	fmt.Printf("平均每个协程启动耗时：%v\n", elapsed/time.Duration(total))


	userInfo := map[string]int{
		"age": 25,
		"socre": 90,
	}
	fmt.Println( userInfo)


	//customized error
	result, err2  :=  utils.Divide2(2,0)
	if err2 != nil{
		fmt.Printf("未知错误：%v \n", err2)
	}
	fmt.Printf("计算结果 :", result)
}