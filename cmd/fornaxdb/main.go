package main

import (
	"github.com/FornaxDB/fornaxdb/logger"
	"github.com/FornaxDB/fornaxdb/storage"
)

func main() {
	l := logger.New()
	config := storage.NewDefaultConfig()
	err := storage.Init(config)
	defer func() {
		err := storage.Close()
		if err != nil {
			l.Error(err.Error(), nil)
		}
	}()
	if err != nil {
		l.Error(err.Error(), nil)
	}
}
