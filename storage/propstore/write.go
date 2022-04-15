package propstore

import (
	"bytes"
	"encoding/binary"
)

// Write writes a new property block to the binary file.
func (p *PropStore) Write(property Property) (ID, error) {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &property)
	if err != nil {
		return -1, err
	}

	var writeSize int
	// TODO: Add validations here to make sure that writeSize == PropertyBlockSize
	writeSize, err = p.File.Write(buffer.Bytes())
	if err != nil {
		return -1, err
	}

	p.Position += int64(writeSize)
	return ID(p.Position) - PropertyBlockSize, nil
}

// Update overwrites the property block at position id with newNode
func (p *PropStore) Update(id ID, newProperty Property) error {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, &newProperty)
	if err != nil {
		return err
	}

	_, err = p.File.WriteAt(buffer.Bytes(), int64(id))
	if err != nil {
		return err
	}

	return nil
}
