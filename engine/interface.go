package engine

type Engine interface {
	Set(key []byte, value interface{}) error
	Get(key []byte) Iterator
	GetRange(left, right []byte) (begin Iterator, end Iterator)
}

type Iterator interface {
	Next() Iterator
	Prev() Iterator
	Key() []byte
	Value() interface{}
}
