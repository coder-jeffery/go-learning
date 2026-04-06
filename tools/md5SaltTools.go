package main


import (
	"crypto/md5"
"encoding/hex"
"fmt"
"math/rand"
"time"
)

// 字符集（生成随机盐用）
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机盐（推荐 16 位）
func generateSalt(length int) string {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, length)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}
	return string(salt)
}

// ========================
// MD5 加盐加密
// ========================
func Md5WithSalt(password string, salt string) string {
	// 组合：密码 + 盐（顺序可自定义）
	str := password + salt
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// ========================
// 密码验证（登录时用）
// ========================
func VerifyPassword(password, salt, encrypted string) bool {
	return Md5WithSalt(password, salt) == encrypted
}

func main() {
	// 1. 用户密码
	password := "123456"

	// 2. 生成随机盐
	salt := generateSalt(16)
	fmt.Println("随机盐 Salt：", salt)

	// 3. 加盐 MD5 加密
	encrypted := Md5WithSalt(password, salt)
	fmt.Println("加盐后 MD5：", encrypted)

	// 4. 验证密码（登录逻辑）
	isOk := VerifyPassword(password, salt, encrypted)
	fmt.Println("密码验证结果：", isOk) // true
}
