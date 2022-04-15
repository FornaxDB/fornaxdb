package heapstore

import (
	"bytes"
	"encoding/binary"
)

// Write at the end of the file
func (h *HeapStore) Write(data []byte) (ID, error) {
	var heapBlocks []HeapBlock

	for i := 0; i < len(data); i += HeapBlockDataSize {
		var hbData [HeapBlockDataSize]byte

		end := i + HeapBlockDataSize
		if end > len(data) {
			end = len(data)
		}

		for j := 0; i+j < end; j++ {
			hbData[j] = data[i+j]
		}

		for j := end; j < HeapBlockDataSize; j++ {
			hbData[j] = 0
		}

		hb := HeapBlock{
			Data: hbData,
			Next: ID(-1),
		}

		heapBlocks = append(heapBlocks, hb)
	}

	// Write to disk in reverse and update the Next prop of the previous HeapBlock
	for i := len(heapBlocks) - 1; i >= 0; i-- {
		var buffer bytes.Buffer
		err := binary.Write(&buffer, binary.BigEndian, heapBlocks[i])
		if err != nil {
			return -1, err
		}

		var writeSize int
		writeSize, err = h.File.Write(buffer.Bytes())
		if err != nil {
			return -1, err
		}

		h.Position += int64(writeSize)

		if i != 0 {
			heapBlocks[i-1].Next = ID(h.Position) - HeapBlockSize
		}
	}

	return ID(h.Position) - HeapBlockSize, nil
}

// Update will update data of a given ID
func (h *HeapStore) Update(id ID, newStr string) error {
	// TODO: Rethink this
	return nil
}
