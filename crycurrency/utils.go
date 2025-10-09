package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 工具函数文件
func uintToBytes(num uint64) []byte {

	//使用binary.Write进行编码
	var buffer bytes.Buffer

	//编码进行错误检查 一定要做
	err := binary.Write(&buffer, binary.BigEndian, num) //对齐参数
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
