package propstore

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

var LOGGER = logger.New()

func (p *PropStore) Init(file *os.File, position int64) {
	p.File = file
	p.Position = position
	LOGGER.Info("Initialised PropStore", nil)
}

func (p *PropStore) Close() error {
	err := p.File.Close()
	if err != nil {
		return err
	}

	LOGGER.Info("Closed PropStore", nil)
	return nil
}
