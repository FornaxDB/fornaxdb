package edgestore

import (
	"bytes"
	"encoding/binary"
)

// Read reads EdgeBlockSize bytes from the id position and marshals into the Edge struct
func (n *EdgeStore) Read(id ID) (*Edge, error) {
	data := make([]byte, EdgeBlockSize)
	_, err := n.File.ReadAt(data, int64(id))
	if err != nil {
		return nil, err
	}

	edge := Edge{}
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &edge)
	if err != nil {
		return nil, err
	}

	return &edge, nil
}
