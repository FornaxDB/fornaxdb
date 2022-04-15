package storage

import (
	"github.com/FornaxDB/fornaxdb/errors"
	"github.com/FornaxDB/fornaxdb/logger"
	"github.com/FornaxDB/fornaxdb/storage/edgestore"
	"github.com/FornaxDB/fornaxdb/storage/heapstore"
	"github.com/FornaxDB/fornaxdb/storage/nodestore"
	"github.com/FornaxDB/fornaxdb/storage/propstore"
	"os"
)

// Config is the configuration for the storage layer
type Config struct {
	Directory string
	Reload    bool
	// TODO: add other relevant config here
}

const (
	NodeStoreFileName   = "fornax.nodestore.db"
	EdgeStoreFileName   = "fornax.edgestore.db"
	PropStoreFileName   = "fornax.propstore.db"
	StringStoreFileName = "fornax.stringstore.db"
	ArrayStoreFileName  = "fornax.arraystore.db"
)

func NewDefaultConfig() *Config {
	return &Config{
		Directory: "/tmp/fornaxdb",
		Reload:    true,
	}
}

var Log logger.Logger
var NodeStore nodestore.NodeStore
var EdgeStore edgestore.EdgeStore
var PropStore propstore.PropStore
var StringStore heapstore.HeapStore
var ArrayStore heapstore.HeapStore

func Init(config *Config) error {
	Log = logger.New()

	var err error
	err = os.MkdirAll(config.Directory, os.ModePerm)
	if err != nil {
		return err
	}

	// TODO: add some metadata to the head of the files, stuff like last updated etc.

	if config.Reload {
		var nodeStoreBinFile *os.File
		nodeStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+NodeStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		// TODO: correct the positions here, these won't be 0 if a file already exists
		NodeStore.Init(nodeStoreBinFile, 0)

		var edgeStoreBinFile *os.File
		edgeStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+EdgeStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		// TODO: correct the positions here, these won't be 0 if a file already exists
		EdgeStore.Init(edgeStoreBinFile, 0)

		var propStoreBinFile *os.File
		propStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+PropStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		// TODO: correct the positions here, these won't be 0 if a file already exists
		PropStore.Init(propStoreBinFile, 0)

		var stringStoreBinFile *os.File
		stringStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+StringStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		// TODO: correct the positions here, these won't be 0 if a file already exists
		StringStore.Init(stringStoreBinFile, 0)

		var arrayStoreBinFile *os.File
		arrayStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+ArrayStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		// TODO: correct the positions here, these won't be 0 if a file already exists
		ArrayStore.Init(arrayStoreBinFile, 0)
	} else {
		var nodeStoreBinFile *os.File
		nodeStoreBinFile, err = os.Create(
			config.Directory + "/" + NodeStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		NodeStore.Init(nodeStoreBinFile, 0)

		var edgeStoreBinFile *os.File
		edgeStoreBinFile, err = os.Create(
			config.Directory + "/" + EdgeStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		EdgeStore.Init(edgeStoreBinFile, 0)

		var propStoreBinFile *os.File
		propStoreBinFile, err = os.Create(
			config.Directory + "/" + PropStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		PropStore.Init(propStoreBinFile, 0)

		var stringStoreBinFile *os.File
		stringStoreBinFile, err = os.Create(
			config.Directory + "/" + StringStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StringStore.Init(stringStoreBinFile, 0)

		var arrayStoreBinFile *os.File
		arrayStoreBinFile, err = os.Create(
			config.Directory + "/" + ArrayStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		ArrayStore.Init(arrayStoreBinFile, 0)
	}

	Log.Info("FornaxDB Storage Initialised Successfully", nil)
	return nil
}

func Close() error {
	var err error

	err = NodeStore.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	err = EdgeStore.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	err = PropStore.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	return nil
}
