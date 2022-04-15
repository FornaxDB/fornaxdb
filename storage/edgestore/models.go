package edgestore

import "os"

type ID int64

type Edge struct {
	Src     ID // ID of Source Node
	Dst     ID // ID of Destination Node
	SrcPrev ID // ID of Previous Edge of Source Node
	SrcNext ID // ID of Next Edge of Source Node
	DstPrev ID // ID of Previous Edge of Destination Node
	DstNext ID // ID of Next Edge of Destination Node
}

type EdgeStore struct {
	File     *os.File
	Position int64
}

const EdgeStrutSize = 48
