package edgestore

import (
	"bytes"
	"encoding/binary"
)

func (n *EdgeStore) Read(id ID) (*Edge, error) {
	// Jump to id position on the file, read an Edge struct
	data := make([]byte, EdgeStrutSize)
	_, err := n.File.ReadAt(data, int64(id))
	if err != nil {
		return nil, err
	}

	node := Edge{}
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}
