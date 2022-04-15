package nodestore

import (
	"bytes"
	"encoding/binary"
)

// Read jumps to the id position in the file and reads NodeBlockSize number of bytes and marshals it into a Node struct
func (n *NodeStore) Read(id ID) (*Node, error) {
	// Jump to id position on the file, read a Node struct
	data := make([]byte, NodeBlockSize)
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
