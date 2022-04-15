package nodestore

import (
	"bytes"
	"encoding/binary"
)

// Write writes a new node at the end of the file
func (n *NodeStore) Write(node Node) (ID, error) {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &node)
	if err != nil {
		return -1, err
	}

	var writeSize int
	writeSize, err = n.File.Write(buffer.Bytes())
	if err != nil {
		return -1, err
	}

	n.Position += int64(writeSize)
	return ID(n.Position - NodeBlockSize), nil
}

// Update will update the node on present on position id
func (n *NodeStore) Update(id ID, newNode Node) error {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &newNode)
	if err != nil {
		return err
	}

	_, err = n.File.WriteAt(buffer.Bytes(), int64(id))
	if err != nil {
		return err
	}

	return nil
}
