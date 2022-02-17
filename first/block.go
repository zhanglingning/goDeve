package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// Block
/*
1、定义结构（区块头的字段比正常的少）
1、前hash
2、当前区块hash
3、数据
*/
type Block struct {
	Version       uint64 // 区块版本号
	PrevBlockHash []byte // 前hash
	MerKleRoot    []byte // 默克尔根，V4使用
	TimeStamp     uint64 // 时间戳
	Difficulity   uint64 // 难度值V2
	Nonce         uint64 // 随机数，挖矿找的就是它
	Data          []byte // 数据
	Hash          []byte // 当前区块hash
}

const genInfo = "hi"

// NewBlock 2、创建区块,对Block对每一个字段填充数据即可
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:       00,
		PrevBlockHash: prevBlockHash,
		MerKleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulity:   1,        // V2
		Nonce:         1,        // V2
		Hash:          []byte{}, // 先填充为空
		Data:          []byte(data),
	}
	block.setHash()
	return &block
}

// 3、生成hash
func (block *Block) setHash() {
	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerKleRoot,
		uintToByte(block.Difficulity),
		uintToByte(block.TimeStamp),
		uintToByte(block.Nonce),
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})

	// 所有的数据，包裹，运算成hash
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
