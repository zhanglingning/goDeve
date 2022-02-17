/*
 * @Author: your name
 * @Date: 2021-08-10 15:05:34
 * @LastEditTime: 2021-12-30 10:08:57
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \blockChain_Fabric_contract1.0d:\GoWork\blockChainV1\block\blockChain\blockchain.go
 */
package main

import "fmt"

// BlockChain 4、创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	// 1、创建一个区块
	// bc.Blocks的最后一个区块的hash值就是当前新区块的PrevBlockHash
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	// 2、添加bc.Blocks数组中
	bc.Blocks = append(bc.Blocks, block)
}

// NewBlockChain //5、引入区块链实现区块链对方法
func NewBlockChain() *BlockChain {
	// 在创建的时候，添加一个区块：创世块
	gensisBlock := NewBlock(genInfo, []byte{0x0000000000000000})
	// 初始化区块链
	bc := BlockChain{
		Blocks: []*Block{
			gensisBlock,
		},
	}
	return &bc
}

func (bc *BlockChain) BlockChainList() {
	for i, block := range bc.Blocks {
		fmt.Printf("I: %d\n", i)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println("====================")
	}
}
