package storage

type StringProp struct {
	Key   string
	Value string
	Next  uint64
}

type IntProp struct {
	Key   string
	Value int
	Next  uint64
}

type FloatProp struct {
	Key   string
	Value float64
	Next  uint64
}

type StringArrProp struct {
	Key   string
	Value []string
	Next  uint64
}

type IntArrProp struct {
	Key   string
	Value []int
	Next  uint64
}

type BoolArrProp struct {
	Key   string
	Value []bool
	Next  uint64
}

type FloatArrProp struct {
	Key   string
	Value []float64
	Next  uint64
}

type ObjProp struct {
	// TODO
}
