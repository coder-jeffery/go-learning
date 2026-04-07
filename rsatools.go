//go:build ignore

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

// ###########################
// 1. 生成 RSA 公私钥对
// ###########################
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
	}
	// 生成公钥
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

// ###########################
// 2. 公钥加密
// ###########################
func rsaEncrypt(plainText string, publicKey *rsa.PublicKey) ([]byte, error) {
	msg := []byte(plainText)
	// RSA 加密（PKCS#1 v1.5 填充模式）
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// ###########################
// 3. 私钥解密
// ###########################
func rsaDecrypt(cipherText []byte, privateKey *rsa.PrivateKey) (string, error) {
	plainBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

func main() {
	// ========== 1. 生成密钥 ==========
	privateKey, publicKey := generateRSAKeyPair(2048) // 2048 位安全强度
	fmt.Println("✅ 密钥生成完成")

	// ========== 2. 待加密原文 ==========
	original := "job@rustdesk.com"
	fmt.Println("原文：", original)

	// ========== 3. 公钥加密 ==========
	cipherText, err := rsaEncrypt(original, publicKey)
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Printf("加密结果（十六进制）：%x\n", cipherText)

	// ========== 4. 私钥解密 ==========
	decryptedText, err := rsaDecrypt(cipherText, privateKey)
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Println("解密结果：", decryptedText)
}
