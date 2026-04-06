package main


import (
	"fmt"
"sync/atomic"
)

// ===================== 配置 =====================
const (
	baseURL       = "http://short.url/"   // 你的短链域名
	base62Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	seq        uint64 = 1 // 自增序列（模拟分布式ID）
	linkMap    = make(map[string]string)
	reverseMap = make(map[string]string)
)

// ===================== Base62 编码（核心）=====================
func toBase62(num uint64) string {
	if num == 0 {
		return string(base62Charset[0])
	}
	var res []byte
	for num > 0 {
		rem := num % 62
		res = append(res, base62Charset[rem])
		num /= 62
	}
	// 反转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// ===================== 生成短链接 =====================
func GenerateShortURL(longURL string) string {
	// 已经生成过，直接返回
	if short, ok := reverseMap[longURL]; ok {
		return baseURL + short
	}

	// 原子自增 ID（并发安全）
	id := atomic.AddUint64(&seq, 1)
	short := toBase62(id)

	// 存储映射
	linkMap[short] = longURL
	reverseMap[longURL] = short

	return baseURL + short
}

// ===================== 解析短链接 =====================
func ResolveShortURL(shortURL string) (string, bool) {
	// 只取最后一段 如 http://short.url/abc → abc
	short := shortURL[len(baseURL):]
	long, ok := linkMap[short]
	return long, ok
}

// ===================== 测试 =====================
func main() {
	long := "https://www.google.com/search?q=golang"

	// 生成短链
	short := GenerateShortURL(long)
	fmt.Println("长链接:", long)
	fmt.Println("短链接:", short)

	// 解析短链
	origin, ok := ResolveShortURL(short)
	if ok {
		fmt.Println("短链还原:", origin)
	} else {
		fmt.Println("短链不存在")
	}
}
