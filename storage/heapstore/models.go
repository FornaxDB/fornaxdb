package heapstore

import "os"

type ID int64

const HeapBlockDataSize = 64 // Bytes
const HeapBlockSize = 64 + 8 // Bytes

// HeapBlock represents a block in the heap file. This block stores 64 bytes of data with a pointer to the next block.
// This pointer is the position of the block in the file.
type HeapBlock struct {
	Data [HeapBlockDataSize]byte
	Next ID
}

// HeapStore is an abstraction over the binary file that holds dynamically sized data in the database
type HeapStore struct {
	File     *os.File
	Position int64
}
