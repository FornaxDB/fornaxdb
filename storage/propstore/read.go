package propstore

import (
	"bytes"
	"encoding/binary"
)

// Read jumps to the id position on the file and reads PropertyBlockSize bytes and marshals that into a Property struct.
func (p *PropStore) Read(id ID) (*Property, error) {
	// Jump to id position on the file, read a Property struct
	data := make([]byte, PropertyBlockSize)
	_, err := p.File.ReadAt(data, int64(id))
	if err != nil {
		return nil, err
	}

	property := Property{}
	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.BigEndian, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}
