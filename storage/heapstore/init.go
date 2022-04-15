package heapstore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

// Init initialises a HeapStore
func (h *HeapStore) Init(file *os.File, position int64) {
	h.File = file
	h.Position = position
	LOGGER.Info("Initialised HeapStore", nil)
}

// Close closes the file held by the HeapStore
func (h *HeapStore) Close() error {
	err := h.File.Close()
	if err != nil {
		return err
	}

	LOGGER.Info("Closing HeapStore", nil)
	return nil
}
