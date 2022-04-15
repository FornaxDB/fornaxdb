package nodestore

import "os"

type Node struct {
	FirstEdge int64
	FirstProp int64
}

type NodeStore struct {
	File     *os.File
	Position int64
}

type ID int64

const NodeStructSize = 16
