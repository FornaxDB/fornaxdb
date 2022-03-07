package storage

type Prop struct {
	Key   string
	Value string
	Next  *Prop
}
