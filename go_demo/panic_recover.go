package main

import "fmt"

//panic_recover

func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b ", err)
		}
	}()
	panic("panic is b")
}

func c() {
	fmt.Println("C")
}

func main() {
	a()
	b()
	c()
}
