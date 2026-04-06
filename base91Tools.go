package main

import (
	"bytes"
	"fmt"
)

// Base91 标准字符集
var base91Table = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '#', '$',
	'%', '&', '(', ')', '*', '+', ',', '.', '/', ':', ';', '<', '=',
	'>', '?', '@', '[', ']', '^', '_', '`', '{', '|', '}', '~', '"',
}

// 解码映射表
var decodeTable [256]int

func init() {
	for i := range decodeTable {
		decodeTable[i] = -1
	}
	for i, c := range base91Table {
		decodeTable[c] = i
	}
}

// Base91 编码（标准、无依赖）
func Base91Encode(data []byte) string {
	var out bytes.Buffer
	var bits uint32
	var bitCount int

	for _, b := range data {
		bits |= uint32(b) << bitCount
		bitCount += 8

		if bitCount >= 13 {
			var n uint32
			if bits&0x1FFF > 88 {
				n = bits & 0x1FFF
				bits >>= 13
				bitCount -= 13
			} else {
				n = bits & 0x3FFF
				bits >>= 14
				bitCount -= 14
			}
			out.WriteByte(base91Table[n%91])
			out.WriteByte(base91Table[n/91])
		}
	}

	if bitCount > 0 {
		out.WriteByte(base91Table[bits%91])
		if bitCount > 7 || bits > 88 {
			out.WriteByte(base91Table[bits/91])
		}
	}

	return out.String()
}

// Base91 解码
func Base91Decode(s string) []byte {
	var out bytes.Buffer
	var bitCount int
	var val uint32

	for _, c := range s {
		d := decodeTable[c]
		if d == -1 {
			continue
		}
		val |= uint32(d) << bitCount
		bitCount += 13

		for bitCount >= 8 {
			out.WriteByte(byte(val & 0xFF))
			val >>= 8
			bitCount -= 8
		}
	}

	if val > 0 {
		out.WriteByte(byte(val))
	}

	return out.Bytes()
}

// 测试入口
func main() {
	original := "job@rustdesk.com Base91 原生实现"

	// 编码
	encoded := Base91Encode([]byte(original))
	fmt.Printf("Base91 编码:", encoded)

	// 解码
	decoded := Base91Decode(encoded)
	fmt.Println("Base91 解码:", string(decoded))
}