package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//1.定义结构 (区块头的字段比正常的少) a.前区块哈希 b.当前区块哈希 c.数据
//2.创建区块
//3.生成哈希
//4.引入区块链
//5.添加区块
//6.重构代码

const genesisInfo = "The Times 03/Jan/2009 Go Programming Language"

//更新区块字段

type Block struct {
	Version       uint64 //区块版本号码
	PrevBlockHash []byte //前区块哈希
	MerkleRoot    []byte //填写为空 后续v4使用
	Timestamp     uint64 //1970-
	Difficulty    uint64
	Nonce         uint64 //随机数
	Hash          []byte //当前区块哈希 区块中本不存在的字段
	Data          []byte //数据 目前使用字节流 v4开始交易使用
}

// 创建区块 对Block的每个字段填充数据
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    []byte{},
		Timestamp:     uint64(time.Now().Unix()),
		Difficulty:    10,
		Nonce:         10,
		Hash:          []byte{}, //填充为空  后续填写数据
		Data:          []byte(data),
	}

	block.SetHash()
	return &block
}

// 生成区块哈希 实现简单的函数 来计算哈希值 无随机值 没有难度值
func (block *Block) SetHash() {
	//var data []byte
	//data = append(data, uintToBytes(block.Version)...)
	//data = append(data, block.PrevBlockHash...)
	//data = append(data, block.MerkleRoot...)
	//data = append(data, uintToBytes(block.Timestamp)...)
	//data = append(data, uintToBytes(block.Difficulty)...)
	//data = append(data, block.Data...)
	//data = append(data, uintToBytes(block.Nonce)...)

	temp := [][]byte{
		uintToBytes(block.Version),
		block.PrevBlockHash,
		block.MerkleRoot,
		uintToBytes(block.Timestamp),
		uintToBytes(block.Difficulty),
		block.Data,
		uintToBytes(block.Nonce),
	}

	data := bytes.Join(temp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// 创建区块链 使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}
