//go:build ignore

package main

import "fmt"

/**
2.搜索旋转排序数组 输入: nums = [4,5,6,7,0,1,2], target = 0 输出target 的数组下标: 4
*/

// search 在旋转排序数组中查找目标值，返回下标，不存在返回-1
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 左半部分有序
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半部分有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(Search(nums, target)) // 输出: 4
}
