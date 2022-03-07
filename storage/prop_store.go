package storage

type Prop[T string | int | bool | float64] struct {
	Key   string
	Value T
	Next  uint64
}

type ArrProp[T []string | []int | []bool | []float64] struct {
	Key   string
	Value T
	Next  uint64
}

type ObjProp struct {
	// TODO
}
