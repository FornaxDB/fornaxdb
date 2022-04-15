package nodestore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

func (n *NodeStore) Init(file *os.File, position int64) {
	n.File = file
	n.Position = position
}

func (n *NodeStore) Close() error {
	err := n.File.Close()
	if err != nil {
		return err
	}

	return nil
}
