package edgestore

import "os"

// ID is the position of the Edge record on disk.
type ID int64

// Edge represents an edge between two nodes in the graph.
type Edge struct {
	Src     ID // ID of Source Node
	Dst     ID // ID of Destination Node
	SrcPrev ID // ID of Previous Edge of Source Node
	SrcNext ID // ID of Next Edge of Source Node
	DstPrev ID // ID of Previous Edge of Destination Node
	DstNext ID // ID of Next Edge of Destination Node
}

// EdgeStore is an abstraction over the binary file that holds all the edges in the database.
type EdgeStore struct {
	File     *os.File
	Position int64
}

// EdgeBlockSize is the total size of the Edge struct in bytes
const EdgeBlockSize = 48 // sizeof(ID) * 6
