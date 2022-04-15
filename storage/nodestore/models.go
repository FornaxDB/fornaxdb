package nodestore

import "os"

// Node represents a node in the database.
type Node struct {
	FirstEdge int64
	FirstProp int64
}

// NodeStore is an abstraction over the binary file that holds the nodes in the database
type NodeStore struct {
	File     *os.File
	Position int64
}

type ID int64

const NodeBlockSize = 16
