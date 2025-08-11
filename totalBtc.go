package main

import (
	"fmt"
)

func main() {

	//fmt.Printf("%x\n", sha256.Sum256([]byte("Hello World")))
	//block := NewBlock(genesisInfo, []byte{0x0000000000000000})

	fmt.Printf("start testing..........................  \n")

	bc := NewBlockChain()
	bc.AddBlock("First Block")

	for _, block := range bc.Blocks {
		fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ \n")
		fmt.Printf("PrevBlockHash value: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash value: %x\n", block.Hash)
		fmt.Printf("Data value: %x\n", block.Data)
	}

	fmt.Printf("testing over....................................\n")

	//data := "hello world"

	//for i := 0; i < 1000000; i++ {
	//	hash := sha256.Sum256([]byte(data + string(i)))
	//	fmt.Printf("hash: %x\n  nonce: %d\n", hash, i)
	//}
}
