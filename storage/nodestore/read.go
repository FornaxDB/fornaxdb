package nodestore

import (
	"bytes"
	"encoding/binary"
)

func (n *NodeStore) Read(id ID) (*Node, error) {
	// Jump to id position on the file, read a Node struct
	data := make([]byte, NodeStructSize)
	_, err := n.File.ReadAt(data, int64(id))
	if err != nil {
		return nil, err
	}

	node := Node{}
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}
