package crc

import (
	"hash/crc32"
)

func CRC(bytes []byte) uint32 {
	hash := crc32.NewIEEE()
	hash.Write(bytes) // feed the bytes into the hash
	return hash.Sum32()
}
