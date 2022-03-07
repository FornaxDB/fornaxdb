package storage

type Edge struct {
	Src     *Edge
	Dst     *Edge
	SrcPrev *Edge
	DstPrev *Edge
	SrcNext *Edge
	DstNext *Edge
}
