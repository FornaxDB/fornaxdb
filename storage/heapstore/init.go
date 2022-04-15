package heapstore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

func (h *HeapStore) Init(file *os.File, position int64) {
	h.File = file
	h.Position = position
	LOGGER.Info("HeapStore Initialised Successfully", nil)
}

func (h *HeapStore) Close() error {
	err := h.File.Close()
	if err != nil {
		return err
	}

	return nil
}
