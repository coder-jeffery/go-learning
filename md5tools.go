package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

//am9ic0BydXN0ZGVzay5jb20=

// 计算字符串的 MD5 值（32位小写）
func md5Encode(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// MD5 破解：通过字典库查表
// 参数：需要破解的md5值、字典字符串（用逗号分隔）
// 返回：原字符串、是否找到
func md5Decode(md5Str string, dict string) (string, bool) {
	// 统一转小写，避免大小写问题
	md5Str = strings.ToLower(md5Str)

	// 分割字典
	words := strings.Split(dict, ",")
	for _, word := range words {
		// 计算字典中每个词的 MD5
		if md5Encode(word) == md5Str {
			return word, true
		}
	}

	return "", false
}

// am9ic0BydXN0ZGVzay5jb20=
func main() {
	// 1. 测试：计算 MD5
	testStr := "123456"
	md5Val := md5Encode(testStr)
	fmt.Println("原字符串：", testStr)
	fmt.Println("MD5值：", md5Val)

	// 2. 测试：破解 MD5
	// 字典：常见密码（可以无限扩充）
	dict := "123456,666666,888888,admin,root,123123,qwerty"

	result, ok := md5Decode(md5Val, dict)
	if ok {
		fmt.Println("破解成功：", result)
	} else {
		fmt.Println("破解失败：字典中未找到对应明文")
	}
}