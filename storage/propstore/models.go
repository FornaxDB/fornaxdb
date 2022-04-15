package propstore

import "os"

type PropType int8

const (
	NullPropType PropType = iota
	BoolPropType
	IntPropType
	FloatPropType
	StringPropType
	ArrayPropType
)

type ID int64

type Property struct {
	Name         ID // ID from string store
	Type         PropType
	IsInlined    bool
	InlinedValue [8]byte
	NextProp     ID
	HeapStorePtr ID
}

type PropStore struct {
	File     *os.File
	Position int64
}

const PropertyBlockSize = 34
