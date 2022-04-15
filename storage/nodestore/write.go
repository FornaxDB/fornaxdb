package nodestore

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Write at the end of the file
func (n *NodeStore) Write(node Node) error {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &node)
	if err != nil {
		return err
	}

	var length int
	length, err = n.File.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	LOGGER.Info(fmt.Sprintf("Wrote %v Bytes", length), nil)
	return nil
}

// Update will update data of a given ID
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
