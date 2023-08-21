package main

import (
	"fmt"
	"hash/crc32"
)

func crc16(data []byte) uint16 {
	crcTable := crc32.MakeTable(0x8005) // CRC-16-IBM

	// 使用CRC32计算得到32位的校验值
	crc := crc32.Checksum(data, crcTable)

	// 取校验值的低16位作为CRC16校验值
	crc16 := uint16(crc & 0xffff)
	return crc16
}

func main() {
	data := []byte("hello world")
	result := crc16(data)
	fmt.Println(result)
	//fmt.Printf("CRC16: 0x%04x\n", result)
}
