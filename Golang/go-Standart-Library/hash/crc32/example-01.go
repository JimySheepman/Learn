package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32q))
}
