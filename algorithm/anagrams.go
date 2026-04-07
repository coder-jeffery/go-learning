package algorithm

import (
	"sort"
	"strings"
)

// Anagrams | golang实现算法
//字母异位词分组

func GroupAnagramsPlus(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int // 字母计数
		for _, c := range str {
			count[c-'a']++ // 统计 a-z 出现次数
		}
		// 计数数组作为 key
		hashMap[count] = append(hashMap[count], str)
	}

	// 转结果
	res := make([][]string, 0)
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}

func GroupAnagrams(strs []string) [][]string {
	// key: 排序后的字符串, value: 对应的异位词列表
	hashMap := make(map[string][]string)

	for _, str := range strs {
		// 1. 把字符串拆分成字符切片
		split := strings.Split(str, "")
		// 2. 排序
		sort.Strings(split)
		// 3. 重新拼接成 key
		key := strings.Join(split, "")

		// 4. 存入 map
		hashMap[key] = append(hashMap[key], str)
	}

	// 5. map values 转成结果二维数组
	res := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}
