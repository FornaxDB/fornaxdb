package edgestore

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (n *EdgeStore) Write(edge Edge) error {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &edge)
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
