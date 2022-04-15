package nodestore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

// Init initialises a NodeStore
func (n *NodeStore) Init(file *os.File, position int64) {
	n.File = file
	n.Position = position

	LOGGER.Info("Initialised NodeStore", nil)
}

// Close closes the file held by the NodeStore
func (n *NodeStore) Close() error {
	err := n.File.Close()
	if err != nil {
		return err
	}

	LOGGER.Info("Closed PropStore", nil)
	return nil
}
