package main


import (
	"encoding/base64"
"fmt"
)

// 编码：字符串 → Base64
func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解码：Base64 → 字符串
func base64Decode(b64Str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// URL 安全编码
func urlSafeBase64Encode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

// URL 安全解码
func urlSafeBase64Decode(b64Str string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(b64Str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	// 测试编码
	original := "job@rustdesk.com"
	encoded := base64Encode(original)
	fmt.Println("编码结果:", encoded) // 输出：am9ic0BydXN0ZGVzay5jb20=

	// 测试解码
	decoded, err := base64Decode(encoded)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}
	fmt.Println("解码结果:", decoded) // 输出：job@rustdesk.com
}
