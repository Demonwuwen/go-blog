package main

import "fmt"

//0 定义区块结构
type Block struct {
	//1. 前区块hash
	PrevHash []byte
	//2. 当前区块hash
	Hash []byte
	//3. 数据
	Data []byte
}

//2。 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash: []byte{},
		Data: []byte(data),
	}
	return &block
}
//3。 生成hash
//4。 引入区块链
//5。 添加区块
//6。 重构代码

func main() {
	block := NewBlock("转一百给我！",[]byte{})
	fmt.Printf("前区块hash：%x\n",block.PrevHash)
	fmt.Printf("当前区块hash：%x\n",block.Hash)
	fmt.Printf("区块data：%s\n",block.Data)
}
