package edgestore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

func (n *EdgeStore) Init(file *os.File, position int64) {
	n.File = file
	n.Position = position
}

func (n *EdgeStore) Close() error {
	err := n.File.Close()
	if err != nil {
		return err
	}

	return nil
}
