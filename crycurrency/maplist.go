package main

import (
	"fmt"
	"strings"
)

func main() {
	//Map 集合
	var a map[string]int = make(map[string]int)
	fmt.Println(a)

	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//判断某个键值对在不在
	var sockerMap = make(map[string]int, 8)
	sockerMap["孙悟空"] = 100
	sockerMap["猪八戒"] = 200
	sockerMap["白龙马"] = 300
	fmt.Println(sockerMap)
	//判断
	value, ok := sockerMap["唐僧"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("唐僧没被抓走", value)
	} else {
		fmt.Println("唐僧又被妖精转走啦")
	}

	fmt.Println("******************************************")
	// for range
	for k, v := range sockerMap {
		fmt.Println(k, v)
	}

	for k := range sockerMap {
		fmt.Println(k)
	}

	for _, v := range sockerMap {
		fmt.Println(v)
	}
	fmt.Println("******************************************")
	//遍历
	delete(sockerMap, "白龙马")
	fmt.Println(sockerMap)

	fmt.Println("******************************************")
	var sliceMap = make(map[string][]int, 8)
	v, ok := sliceMap["china"]
	if ok {
		fmt.Println(v)
	} else {
		//sliceMap初始化
		sliceMap["china"] = make([]int, 3) //init finish
		sliceMap["china"][0] = 100
		sliceMap["china"][1] = 200
		sliceMap["china"][2] = 300
		//fmt.Println(sliceMap["china"])

		for k, v := range sliceMap {
			fmt.Println(k, v)
		}
	}

	fmt.Println("******************************************")
	//map统计
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
