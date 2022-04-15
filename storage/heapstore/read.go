package heapstore

import (
	"bytes"
	"encoding/binary"
)

func (h *HeapStore) Read(id ID) ([]byte, error) {
	var result []byte
	currentId := id

	for currentId != -1 {
		data := make([]byte, HeapBlockSize)
		_, err := h.File.ReadAt(data, int64(currentId))
		if err != nil {
			return nil, err
		}

		heapBlock := HeapBlock{}
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.BigEndian, &heapBlock)
		if err != nil {
			return nil, err
		}

		for _, b := range heapBlock.Data {
			result = append(result, b)
		}

		currentId = heapBlock.Next
	}

	return result, nil
}
