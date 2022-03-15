package storage

import (
	"github.com/FornaxDB/fornaxdb/errors"
	"github.com/FornaxDB/fornaxdb/logger"
	"os"
)

// Config is the configuration for the storage layer
type Config struct {
	Directory string
	Reload    bool
	// TODO: add other relevant config here
}

type StoreFile struct {
	File     *os.File
	Position uint64
}

// Files has the file pointers to all the binary files for the graph data
type Files struct {
	NodeStoreFile *StoreFile
	EdgeStoreFile *StoreFile
	PropStoreFile *StoreFile
}

type ID uint64

const (
	NodeStoreFileName = "fornax.nodestore.db"
	EdgeStoreFileName = "fornax.edgestore.db"
	PropStoreFileName = "fornax.propstore.db"
)

func NewDefaultConfig() *Config {
	return &Config{
		Directory: "/tmp/fornaxdb",
		Reload:    true,
	}
}

var StateFiles Files
var Log logger.Logger

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
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.NodeStoreFile = &StoreFile{
			File:     nodeStoreBinFile,
			Position: 0, // TODO: correct the positions here, these won't be 0 if a file already exists
		}

		var edgeStoreBinFile *os.File
		edgeStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+NodeStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.EdgeStoreFile = &StoreFile{
			File:     edgeStoreBinFile,
			Position: 0,
		}

		var propStoreBinFile *os.File
		propStoreBinFile, err = os.OpenFile(
			config.Directory+"/"+NodeStoreFileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.PropStoreFile = &StoreFile{
			File:     propStoreBinFile,
			Position: 0,
		}
	} else {
		var nodeStoreBinFile *os.File
		nodeStoreBinFile, err = os.Create(
			config.Directory + "/" + NodeStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.NodeStoreFile = &StoreFile{
			File:     nodeStoreBinFile,
			Position: 0,
		}

		var edgeStoreBinFile *os.File
		edgeStoreBinFile, err = os.Create(
			config.Directory + "/" + NodeStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.EdgeStoreFile = &StoreFile{
			File:     edgeStoreBinFile,
			Position: 0,
		}

		var propStoreBinFile *os.File
		propStoreBinFile, err = os.Create(
			config.Directory + "/" + NodeStoreFileName,
		)
		if err != nil {
			return errors.StorageCannotOpenFile.New(err.Error())
		}

		StateFiles.PropStoreFile = &StoreFile{
			File:     propStoreBinFile,
			Position: 0,
		}
	}

	Log.Info("FornaxDB Storage Initialised Successfully", nil)
	return nil
}

func Close() error {
	var err error

	err = StateFiles.NodeStoreFile.File.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	err = StateFiles.EdgeStoreFile.File.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	err = StateFiles.PropStoreFile.File.Close()
	if err != nil {
		return errors.StorageCannotCloseFile.New(err.Error())
	}

	return nil
}
