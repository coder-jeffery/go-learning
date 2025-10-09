//package main
//
//// 实现创建区块链方法
//func NewBlockChain() *main.BlockChain {
//	genesisBlock := NewBlock(genesisInfo, []byte{0x00000000000000000})
//	bc := main.BlockChain{[]*main.Block{genesisBlock}}
//	return &bc
//}
//
//// 添加区块
//func (bc *main.BlockChain) AddBlock(data string) {
//	lastBlock := bc.Blocks[len(bc.Blocks)-1]
//	prevHash := lastBlock.Hash
//	block := NewBlock(data, prevHash)
//	bc.Blocks = append(bc.Blocks, block)
//}
