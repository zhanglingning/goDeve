package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 工具函数文件
func uintToByte(num uint64) []byte {
	//TODO uint 转 byte
	// 使用binary.Write 来进行编码
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
