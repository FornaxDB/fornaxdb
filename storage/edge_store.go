package storage

type Edge struct {
	Src     uint64
	Dst     uint64
	SrcPrev uint64
	DstPrev uint64
	SrcNext uint64
	DstNext uint64
}
