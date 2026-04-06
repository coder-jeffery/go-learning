package main



import (
	"bytes"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"encoding/base64"
"fmt"
"io"
)

// ========================
// 填充函数：让数据长度满足 AES 块要求（16字节）
// ========================
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// ========================
// 去填充函数：解密后去掉填充
// ========================
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

// ========================
// AES 加密（返回 base64 字符串）
// key：密钥（必须 16/24/32 字节，对应 AES-128/192/256）
// ========================
func AESEncrypt(origData string, key []byte) (string, error) {
	data := []byte(origData)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)

	// 生成随机 IV（初始向量），长度=块大小
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// CBC 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	// 把 IV + 密文拼接，再 base64 输出
	result := append(iv, crypted...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// ========================
// AES 解密（传入 base64 加密串）
// ========================
func AESDecrypt(cryptedStr string, key []byte) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(cryptedStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(crypted) < blockSize {
		return "", fmt.Errorf("密文长度错误")
	}

	// 拆分 IV 和 密文
	iv := crypted[:blockSize]
	crypted = crypted[blockSize:]

	// CBC 解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	// 去填充
	origData = pkcs7UnPadding(origData)
	return string(origData), nil
}

func main() {
	// 1. 密钥（16字节 = AES-128，24=192，32=256）
	key := []byte("1234567890123456") // 必须16/24/32位

	// 2. 待加密字符串（你之前的邮箱）
	original := "job@rustdesk.com"
	fmt.Println("原文：", original)

	// 3. 加密
	encryptStr, err := AESEncrypt(original, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("AES 加密结果(base64)：", encryptStr)

	// 4. 解密
	decryptStr, err := AESDecrypt(encryptStr, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("AES 解密结果：", decryptStr)
}
