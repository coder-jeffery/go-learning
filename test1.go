//go:build ignore

package main

import "fmt"


/**
1.只出现一次的数字某个元素只出现一次以外，其余每个元素均出现两次，找出1次的输入：nums = [2,2,1]输出：1
*/

func SingleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(SingleNumber(nums)) // 输出: 1


}
