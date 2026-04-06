package main


import (
"encoding/base64"
"fmt"
)

// ========================
// Base64URL 编码（无填充 = 推荐）
// ========================
func Base64URLEncode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// ========================
// Base64URL 解码
// ========================
func Base64URLDecode(encoded string) (string, error) {
	bytes, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	// 测试字符串（带特殊字符，URL 友好）
	original := "job@rustdesk.com+123/abc"

	// 1. 编码
	encoded := Base64URLEncode(original)
	fmt.Println("Base64URL 编码结果：", encoded)
	// 输出：am9ic0BydXN0ZGVzay5jb20rMTIzL2Ficw

	// 2. 解码
	decoded, err := Base64URLDecode(encoded)
	if err != nil {
		fmt.Println("解码失败：", err)
		return
	}
	fmt.Println("Base64URL 解码结果：", decoded)
	// 输出：job@rustdesk.com+123/abc
}
