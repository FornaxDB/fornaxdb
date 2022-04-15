package heapstore

import "os"

type ID int64

const HeapBlockDataSize = 64 // Bytes
const HeapBlockSize = 64 + 8 // Bytes

type HeapBlock struct {
	Data [HeapBlockDataSize]byte
	Next ID
}

type HeapStore struct {
	File     *os.File
	Position int64
}
