package edgestore

import (
	"bytes"
	"encoding/binary"
)

// Write writes the passed edge to the disk and returns the position of the written record.
func (n *EdgeStore) Write(edge Edge) (ID, error) {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &edge)
	if err != nil {
		return -1, err
	}

	var writeSize int
	writeSize, err = n.File.Write(buffer.Bytes())
	if err != nil {
		return -1, err
	}

	n.Position += int64(writeSize)
	return ID(n.Position - EdgeBlockSize), nil
}

// Update will overwrite the edge present at position id on the disk.
func (n *EdgeStore) Update(id ID, newEdge Edge) error {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &newEdge)
	if err != nil {
		return err
	}

	_, err = n.File.WriteAt(buffer.Bytes(), int64(id))
	if err != nil {
		return err
	}

	return nil
}
