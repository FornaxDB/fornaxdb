package edgestore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

// Init initialises an EdgeStore
func (n *EdgeStore) Init(file *os.File, position int64) {
	n.File = file
	n.Position = position
	LOGGER.Info("Initialized EdgeStore", nil)
}

// Close closes the file held by the EdgeStore
func (n *EdgeStore) Close() error {
	err := n.File.Close()
	if err != nil {
		return err
	}

	LOGGER.Info("Closing EdgeStore", nil)
	return nil
}
